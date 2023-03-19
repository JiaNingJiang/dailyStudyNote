package xclient

import (
	"context"
	"geeRpc/geerpc"
	"io"
	"reflect"
	"sync"
)

type XClient struct {
	d       Discovery  // XClient能够进行服务发现
	mode    SelectMode // 负载均衡策略
	opt     *geerpc.Option
	mu      sync.Mutex                // protect following
	clients map[string]*geerpc.Client // 存储所有已注册的RPC Client，进行复用
}

var _ io.Closer = (*XClient)(nil)

func NewXClient(d Discovery, mode SelectMode, opt *geerpc.Option) *XClient {
	return &XClient{d: d, mode: mode, opt: opt, clients: make(map[string]*geerpc.Client)}
}

// 关闭所有已注册的RPC Client
func (xc *XClient) Close() error {
	xc.mu.Lock()
	defer xc.mu.Unlock()
	for key, client := range xc.clients {
		// I have no idea how to deal with error, just ignore it.
		_ = client.Close()
		delete(xc.clients, key)
	}
	return nil
}

// 参数为RPC Server的连接类型和IP地址
func (xc *XClient) dial(rpcAddr string) (*geerpc.Client, error) {
	xc.mu.Lock()
	defer xc.mu.Unlock()
	client, ok := xc.clients[rpcAddr]
	if ok && !client.IsAvailable() { //与对应RPC Server连接的RPC Client存在但是连接不可用，直接关闭并删除记录
		_ = client.Close()
		delete(xc.clients, rpcAddr)
		client = nil
	}
	if client == nil { // RPC Client 不存在,创建一个新的(根据RPC Server地址信息和Option协商消息)
		var err error
		client, err = geerpc.XDial(rpcAddr, xc.opt) // 创建并返回新的RPC Client
		if err != nil {
			return nil, err
		}
		xc.clients[rpcAddr] = client
	}
	return client, nil // 如果RPC Client存在且可用,则会直接返回
}

// 通过RPC Client 远程调用指定RPC Server上的RPC服务
func (xc *XClient) call(rpcAddr string, ctx context.Context, serviceMethod string, args, reply interface{}) error {
	client, err := xc.dial(rpcAddr)
	if err != nil {
		return err
	}
	return client.Call(ctx, serviceMethod, args, reply)
}

// 负载均衡 + 远程调用RPC服务
func (xc *XClient) Call(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	rpcAddr, err := xc.d.Get(xc.mode) // 根据指定的负载均衡策略,选择一个RPC Server
	if err != nil {
		return err
	}
	return xc.call(rpcAddr, ctx, serviceMethod, args, reply)
}

func (xc *XClient) Broadcast(ctx context.Context, serviceMethod string, args, reply interface{}) error {
	servers, err := xc.d.GetAll()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	var mu sync.Mutex // 并发情况下保护 e 和 replyDone
	var e error
	replyDone := reply == nil // if reply is nil, don't need to set value
	ctx, cancel := context.WithCancel(ctx)
	for _, rpcAddr := range servers { // 向所有能够提供RPC Server的实例请求RPC服务（为了提升性能，请求是并发的）
		wg.Add(1)
		addr := rpcAddr
		go func(rpcAddr string) {
			defer wg.Done()
			var clonedReply interface{}
			if reply != nil {
				clonedReply = reflect.New(reflect.ValueOf(reply).Elem().Type()).Interface()
			}
			err := xc.call(rpcAddr, ctx, serviceMethod, args, clonedReply) // 远程调用指定RPC Server上的RPC服务
			mu.Lock()
			if err != nil && e == nil { // 调用RPC服务失败
				e = err
				cancel() // if any call failed, cancel unfinished calls
			}
			if err == nil && !replyDone { // 仅需要一个服务实例完成本次RPC服务请求即可(replyDone成为true之后不会被再赋值为false)
				reflect.ValueOf(reply).Elem().Set(reflect.ValueOf(clonedReply).Elem()) // reply从clonedReply获取RPC请求的响应参数
				replyDone = true
			}
			mu.Unlock()
		}(addr)
	}
	wg.Wait()
	return e
}
