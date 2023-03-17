package main

import (
	"fmt"
	"geeRpc/geerpc"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string) {
	// pick a free port
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal("network error:", err)
	}
	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()
	geerpc.Accept(l)
}

func main() {
	addr := make(chan string)
	go startServer(addr) // 启动RPC服务端(通过addr管道返回server端的监听地址)

	client, _ := geerpc.Dial("tcp", <-addr) //创建与RPC Server连接的RPC Client
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { // 发送 5次 RPC请求 (使用了等待组和同步方式的client.Call确保RPC请求是串行执行的)
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			args := fmt.Sprintf("geerpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatal("call Foo.Sum error:", err)
			}
			log.Println("reply:", reply)
		}(i)
	}
	wg.Wait() // 所有的RPC请求都完成之后才能退出
}
