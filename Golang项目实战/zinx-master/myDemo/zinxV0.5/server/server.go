package main

import (
	"GO_Demo/zinx/ziface"
	"GO_Demo/zinx/znet"
	"fmt"
)

//服务器自定义的路由0
type pingRouter struct {
	znet.BaseRouter //继承基类
}

//服务器自定义的路由1
type helloRouter struct {
	znet.BaseRouter //继承基类
}

//重写三个路由方法
/*func (r *pingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("[server] Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("[server] call back before ping error...")
	}
}*/

func (r *pingRouter) Handle(request ziface.IRequest) {
	fmt.Println("[server] Call Router Handle...")
	//读取客户端数据，然后回复ping...ping...ping
	fmt.Printf("[server] recv msg from client , msgID: %v ,msg: %v\n", request.GetMsgID(), string(request.GetData()))

	if err := request.GetConnection().Send(100, []byte("ping...ping...ping")); err != nil { //调用send方法，封包之后发送
		fmt.Println("[server] router Handle is err:", err)
	}
}

func (r *helloRouter) Handle(request ziface.IRequest) {
	fmt.Println("[server] Call Router Handle...")
	//读取客户端数据，然后回复ping...ping...ping
	fmt.Printf("[server] recv msg from client , msgID: %v ,msg: %v\n", request.GetMsgID(), string(request.GetData()))

	if err := request.GetConnection().Send(200, []byte("hello...hello...hello")); err != nil { //调用send方法，封包之后发送
		fmt.Println("[server] router Handle is err:", err)
	}
}

/*
func (r *pingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("[server] Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("[server] call back after ping error...")
	}
}*/

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("[server] DoConnectionBegin is called ~~~")
	if err := conn.Send(100, []byte("Pre Connection Hook")); err != nil {
		fmt.Println("[server] DoConnectionBegin is error: ", err)
	}
}

func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("[server] DoConnectionLost is called ~~~")
}

func main() {
	//1.创建一个基于zinx框架的服务器模块
	s := znet.NewServer("[zinxV0.1]")

	//2.注册连接建立后以及断开前需要执行的Hook函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//3.为对于消息ID的消息注册路由处理业务
	s.AddRouter(0, &pingRouter{})
	s.AddRouter(1, &helloRouter{})
	//4.运行服务器
	s.Run()
}
