package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geeRpc/codec"
	"log"
	"net"
	"sync"
)

var CallChanMax int = 10

// Call 包含了一次 RPC 调用所需要的信息。
type Call struct {
	Seq           uint64
	ServiceMethod string      // format "<service>.<method>"
	Args          interface{} // Client发送RPC请求携带的函数参数
	Reply         interface{} // 接收来自RPC服务端的响应 body
	Error         error       // if error occurs, it will be set
	Done          chan *Call  // 支持异步调用：当调用结束时，会调用 call.done() 通知调用方。
}

func (call *Call) done() {
	call.Done <- call
}

type Client struct {
	cc       codec.Codec // 发送和接收编码的RPC请求和应答
	opt      *Option     // 协商请求
	sending  sync.Mutex  // 保证请求的有序发送，即防止出现多个请求报文混淆
	header   codec.Header
	mu       sync.Mutex       // 包含Client客户端的状态量
	seq      uint64           // 用于给发送的请求编号，每个请求拥有唯一编号
	pending  map[uint64]*Call //  存储未处理完的请求
	closing  bool             // 用户主动关闭RPC连接(Client将不可用)
	shutdown bool             //  shutdown 置为 true 一般是有错误发生(Client将不可用)
}

var ErrShutdown = errors.New("connection is shut down")

// 主动请求关闭与RPC服务端的连接
func (client *Client) Close() error {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing {
		return ErrShutdown
	}
	client.closing = true
	return client.cc.Close()
}

// 判断与RPC服务端的连接是否仍可用(true == 可用)
func (client *Client) IsAvailable() bool {
	client.mu.Lock()
	defer client.mu.Unlock()
	return !client.shutdown && !client.closing
}

// 在RPC Client端生成(注册)一个RPC请求(同时按照顺序为其添加序列号)，将该请求添加到执行队列Client.pending中
func (client *Client) registerCall(call *Call) (uint64, error) {
	client.mu.Lock()
	defer client.mu.Unlock()
	if client.closing || client.shutdown {
		return 0, ErrShutdown
	}
	call.Seq = client.seq
	client.pending[call.Seq] = call
	client.seq++
	return call.Seq, nil
}

// 将指定序列号的RPC请求从执行队列Client.pending中删除
func (client *Client) removeCall(seq uint64) *Call {
	client.mu.Lock()
	defer client.mu.Unlock()
	call := client.pending[seq]
	delete(client.pending, seq)
	return call
}

// 服务端或客户端发生错误时调用。强行终止执行队列Client.pending中所有的RPC请求，调用call.done通知调用方RPC请求已结束，同时在call.Error中标注错误信息
func (client *Client) terminateCalls(err error) {
	client.sending.Lock()
	defer client.sending.Unlock()
	client.mu.Lock()
	defer client.mu.Unlock()
	client.shutdown = true
	for _, call := range client.pending {
		call.Error = err
		call.done()
	}
}

// 循环(知道接收发生错误才退出循环)接收来自于RPC服务端的响应
func (client *Client) receive() {
	var err error
	for err == nil {
		var h codec.Header
		if err = client.cc.ReadHeader(&h); err != nil { // 读取并解码响应的消息头
			break // 如果消息头解码错误,则直接退出循环
		}
		call := client.removeCall(h.Seq) // RPC请求已经被服务端处理并得到了响应,因此需要将对应的Call请求从执行队列Client.pending中删除
		switch {
		case call == nil: // 情况一：被服务器执行的RPC请求在本地已经不存在(可能是请求没有发送完整，或者因为其他原因被取消，但是服务端仍旧处理了)
			err = client.cc.ReadBody(nil) // 无视本次RPC服务端的响应(但仍然需要读取,因为响应已经被存入当前客户端的接收缓存中，如不读取则会影响下一次RPC请求的接收)
		case h.Error != "": // 情况二：服务端在处理RPC请求时发生了错误，也意味着本次的RPC响应是无效的
			call.Error = fmt.Errorf(h.Error)
			err = client.cc.ReadBody(nil)
			call.done() // 需要将对应的RPC请求从执行队列Client.pending中删除
		default: // 情况三：正常收发且处理
			err = client.cc.ReadBody(call.Reply)
			if err != nil {
				call.Error = errors.New("reading body " + err.Error())
			}
			call.done()
		}
	}
	client.terminateCalls(err) // RPC服务发生了解码错误(系统级)，对Client.pending中的所有RPC请求实行 关闭并通知错误信息 的操作
}

// 根据与RPC服务端已有的conn socket，向服务端发送一条JSON格式的Option协商消息，同时创建一个新的RPC客户端并返回
func NewClient(conn net.Conn, opt *Option) (*Client, error) {
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		err := fmt.Errorf("invalid codec type %s", opt.CodecType)
		log.Println("rpc client: codec error:", err)
		return nil, err
	}
	// send options with server
	if err := json.NewEncoder(conn).Encode(opt); err != nil { // 将JSON编码后的Option协商消息发送给RPC服务端
		log.Println("rpc client: options error: ", err)
		_ = conn.Close()
		return nil, err
	}
	return newClientCodec(f(conn), opt), nil // 创建一个新的RPC客户端(构造的编解码收发器 和 Option协商消息 作为RPC Client的参数)
}

// 创建并返回一个RPC Client，并由协程负责运行与RPC Server的接收服务
func newClientCodec(cc codec.Codec, opt *Option) *Client {
	client := &Client{
		seq:     1, // seq总是从1开始 (seq = 0/1 都表示无效的RPC请求)
		cc:      cc,
		opt:     opt,
		pending: make(map[uint64]*Call),
	}
	go client.receive() // 协程运行RPC Client接收服务
	return client
}

// 解析传入的 Option协商消息 列表，返回第一个合法的 Option协商消息
func parseOptions(opts ...*Option) (*Option, error) {
	// if opts is nil or pass nil as parameter
	if len(opts) == 0 || opts[0] == nil {
		return DefaultOption, nil
	}
	if len(opts) != 1 {
		return nil, errors.New("number of options is more than 1")
	}
	opt := opts[0]
	opt.MagicNumber = DefaultOption.MagicNumber
	if opt.CodecType == "" {
		opt.CodecType = DefaultOption.CodecType
	}
	return opt, nil
}

// 与RPC Server端建立网络连接，并在此基础之上创建 RPC Client并返回
func Dial(network, address string, opts ...*Option) (client *Client, err error) {
	opt, err := parseOptions(opts...) //
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial(network, address) // 与RPC Service建立通信，返回conn socket
	if err != nil {
		return nil, err
	}
	// close the connection if client is nil
	defer func() {
		if client == nil { // defer机制，在调用defer之前和return之后会完成对返回值的赋值，因此此处调用client其不会为空
			_ = conn.Close()
		}
	}()
	return NewClient(conn, opt)
}

// 向RPC Server 发送 RPC请求(Call)
func (client *Client) send(call *Call) {
	// make sure that the client will send a complete request
	client.sending.Lock()
	defer client.sending.Unlock()

	// register this call.
	seq, err := client.registerCall(call) // 将call在当前RPC Client上完成注册(为call更新seq并添加到执行队列Client.pending中)
	if err != nil {
		call.Error = err
		call.done()
		return
	}

	// prepare request header
	client.header.ServiceMethod = call.ServiceMethod
	client.header.Seq = seq
	client.header.Error = ""

	// encode and send the request
	if err := client.cc.Write(&client.header, call.Args); err != nil { // 向服务端发送RPC请求(包括请求头和服务的函数参数)
		call := client.removeCall(seq) // 如果发送过程产生了错误，则将该请求从执行队列Client.pending中删除
		if call != nil {
			call.Error = err // 获取错误信息
			call.done()      // 通知调用方本次RPC请求结束
		}
	}
}

// 收发分离(异步化)。生成一个RPC请求Call，执行client.send(call)完成发送。同时最后会将本次产生并发送的Call进行返回(方便调用者从Call.Done中获知本次RPC请求已经完成的通知)
func (client *Client) Go(serviceMethod string, args, reply interface{}, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call, CallChanMax) // done 是一个可以重复使用的管道(多个RPC请求进行共享)，因此done的管道容量是可以根据网络流量大小进行调整的
	} else if cap(done) == 0 {
		log.Panic("rpc client: done channel is unbuffered")
	}
	call := &Call{
		ServiceMethod: serviceMethod,
		Args:          args,
		Reply:         reply,
		Done:          done,
	}
	go client.send(call) // 无需等待发送完成，Call.Done通常是在接收到RPC Server的回复是才会接收到信号
	return call
}

// 收发同步。同一个RPC请求Call的发送与接收是同步的，RPC Client在完成RPC请求的发送之后，必须等到Call.Done中接收到信号(意味着本次RPC请求被完成)才会退出本方法
func (client *Client) Call(serviceMethod string, args, reply interface{}) error {
	call := <-client.Go(serviceMethod, args, reply, make(chan *Call, 1)).Done
	return call.Error
}
