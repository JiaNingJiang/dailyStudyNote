package main

import (
	"GO_Demo/zinx/ziface"
	"GO_Demo/zinx/znet"
	"fmt"
)

//服务器自定义的路由
type pingRouter struct {
	znet.BaseRouter //继承基类
}

//重写三个路由方法
func (r *pingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back before ping error...")
	}
}

func (r *pingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("ping...\n"))
	if err != nil {
		fmt.Println("call back ping error...")
	}
}

func (r *pingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTCPConn().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back after ping error...")
	}
}

func main() {
	//创建一个基于zinx框架的服务器模块
	s := znet.NewServer("[zinxV0.1]")

	s.AddRouter(&pingRouter{}) //注册路由处理业务
	//运行服务器
	s.Run()
}
