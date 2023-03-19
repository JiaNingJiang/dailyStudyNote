package main

import (
	"context"
	"geeRpc/geerpc"
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

func startServer(addr chan string) {

	var foo Foo
	if err := geerpc.Register(&foo); err != nil {
		log.Fatal("register error:", err)
	}

	l, err := net.Listen("tcp", ":8888") // ":0"的作用就是随机寻找一个空闲的端口进行监听
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	//geerpc.Accept(l)
	geerpc.HandleHTTP() // 注册http路由

	_ = http.Serve(l, nil) // 提供http服务
}

func call(addrCh chan string) {
	client, _ := geerpc.DialHTTP("tcp", <-addrCh)
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)
	// send request & receive response
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { // 发送 5次 RPC请求 (使用了等待组和同步方式的client.Call确保RPC请求是串行执行的)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := &Args{Num1: i, Num2: i * i}
			var reply int
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second) // RPC Client进行RPC调用时引入了超时检测机制
			if err := client.Call(ctx, "Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Printf("%d + %d = %d", args.Num1, args.Num2, reply)
		}(i)
	}
	wg.Wait() // 所有的RPC请求都完成之后才能退出
}

func main() {
	ch := make(chan string)
	go call(ch)
	startServer(ch) // 启动RPC服务端(通过addr管道返回server端的监听地址)
}
