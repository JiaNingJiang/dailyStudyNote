package main

import "GO_Demo/zinx/znet"

func main() {
	//创建一个基于zinx框架的服务器模块
	s := znet.NewServer("[zinxV0.1]")
	//运行服务器
	s.Run()
}
