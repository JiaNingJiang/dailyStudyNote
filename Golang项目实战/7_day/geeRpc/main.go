package main

import (
	"context"
	"geeRpc/geerpc"
	"geeRpc/registry"
	"geeRpc/xclient"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func (f Foo) Sleep(args Args, reply *int) error {
	time.Sleep(time.Second * time.Duration(args.Num1))
	*reply = args.Num1 + args.Num2
	return nil
}

// 注册并运行服务中心
func startRegistry(wg *sync.WaitGroup) {
	l, _ := net.Listen("tcp", ":9999")
	registry.HandleHTTP()
	wg.Done()
	_ = http.Serve(l, nil)
}

// 注册并运行服务端(无需向客户端暴露自身addr)
func startServer(registryAddr string, wg *sync.WaitGroup) {
	var foo Foo
	l, _ := net.Listen("tcp", ":0")
	server := geerpc.NewServer()
	_ = server.Register(&foo)
	registry.Heartbeat(registryAddr, "tcp@"+l.Addr().String(), 0) // 向注册中心持续发送心跳包
	wg.Done()
	server.Accept(l)
}

// func startServer(addr chan string) {

// 	var foo Foo
// 	server := geerpc.NewServer()
// 	_ = server.Register(&foo)

// 	l, err := net.Listen("tcp", ":0") // ":0"的作用就是随机寻找一个空闲的端口进行监听
// 	if err != nil {
// 		log.Fatal("network error:", err)
// 	}
// 	log.Println("start rpc server on", l.Addr())

// 	addr <- l.Addr().String()
// 	server.Accept(l)

// 	//geerpc.HandleHTTP() // 注册http路由
// 	//_ = http.Serve(l, nil) // 提供http服务
// }

// 封装一个方法 foo，便于在 Call 或 Broadcast 之后统一打印成功或失败的日志。
func foo(xc *xclient.XClient, ctx context.Context, typ, serviceMethod string, args *Args) {
	var reply int
	var err error
	switch typ {
	case "call":
		err = xc.Call(ctx, serviceMethod, args, &reply)
	case "broadcast":
		err = xc.Broadcast(ctx, serviceMethod, args, &reply)
	}
	if err != nil {
		log.Printf("%s %s error: %v", typ, serviceMethod, err)
	} else {
		log.Printf("%s %s success: %d + %d = %d", typ, serviceMethod, args.Num1, args.Num2, reply)
	}
}

// func call(addrCh chan string) {
// 	client, _ := geerpc.XDial("http@" + <-addrCh)

// 	defer func() { _ = client.Close() }()

// 	time.Sleep(time.Second)
// 	// send request & receive response
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ { // 发送 5次 RPC请求 (使用了等待组和同步方式的client.Call确保RPC请求是串行执行的)
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			args := &Args{Num1: i, Num2: i * i}
// 			var reply int
// 			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // RPC Client进行RPC调用时引入了超时检测机制
// 			if err := client.Call(ctx, "Foo.Sum", args, &reply); err != nil {
// 				log.Fatal("call Foo.Sum error:", err)
// 			}
// 			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
// 		}(i)
// 	}
// 	wg.Wait() // 所有的RPC请求都完成之后才能退出
// }

func call(registry string) {
	d := xclient.NewGeeRegistryDiscovery(registry, 0)      // 客户端的服务发现(向注册中心请求服务实例)
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil) // 完成服务发现,返回RPC Client
	defer func() { _ = xc.Close() }()
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // RPC Client进行RPC调用时引入了超时检测机制
			foo(xc, ctx, "call", "Foo.Sum", &Args{Num1: i, Num2: i * i})
		}(i)
	}
	wg.Wait() // 所有的RPC请求都完成之后才能退出
}

// func xcall(addr1, addr2 string) {
// 	d := xclient.NewMultiServerDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2}) // 存在两个提供RPC服务的实例
// 	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
// 	defer func() { _ = xc.Close() }()
// 	// send request & receive response
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
// 			foo(xc, ctx, "call", "Foo.Sum", &Args{Num1: i, Num2: i * i})
// 		}(i)
// 	}
// 	wg.Wait()
// }

func broadcast(registry string) {
	d := xclient.NewGeeRegistryDiscovery(registry, 0)
	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
	defer func() { _ = xc.Close() }()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
			foo(xc, ctx, "broadcast", "Foo.Sum", &Args{Num1: i, Num2: i * i})
			foo(xc, ctx, "broadcast", "Foo.Sleep", &Args{Num1: i, Num2: i * i})
		}(i)
	}
	wg.Wait()
}

// func broadcast(addr1, addr2 string) {
// 	d := xclient.NewMultiServerDiscovery([]string{"tcp@" + addr1, "tcp@" + addr2})
// 	xc := xclient.NewXClient(d, xclient.RandomSelect, nil)
// 	defer func() { _ = xc.Close() }()
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(i int) {
// 			defer wg.Done()
// 			ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
// 			foo(xc, ctx, "broadcast", "Foo.Sum", &Args{Num1: i, Num2: i * i})
// 			// expect 2 - 5 timeout
// 			foo(xc, ctx, "broadcast", "Foo.Sleep", &Args{Num1: i, Num2: i * i})
// 		}(i)
// 	}
// 	wg.Wait()
// }

// func main() {

// 	ch1 := make(chan string)
// 	ch2 := make(chan string)
// 	// start two servers
// 	go startServer(ch1)
// 	go startServer(ch2)

// 	addr1 := <-ch1
// 	addr2 := <-ch2

// 	time.Sleep(time.Second)
// 	//call(ch1)
// 	xcall(addr1, addr2)
// 	broadcast(addr1, addr2)
// }

func main() {
	registryAddr := "http://localhost:9999/_geerpc_/registry"
	var wg sync.WaitGroup
	wg.Add(1)
	go startRegistry(&wg) // 启动并等待注册中心完成启动
	wg.Wait()

	time.Sleep(time.Second)
	wg.Add(2)
	go startServer(registryAddr, &wg) // 启动并等待两个服务端完成启动
	go startServer(registryAddr, &wg)
	wg.Wait()

	time.Sleep(time.Second)
	call(registryAddr)      // 单点调用RPC
	broadcast(registryAddr) // 广播式调用RPC

}
