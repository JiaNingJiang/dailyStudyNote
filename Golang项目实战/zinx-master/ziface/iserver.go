package ziface

type IServer interface {

	//启动服务器
	Start()
	//关闭服务器
	Stop()
	//运行服务器
	Run()
	//路由功能，为当前类型的消息注册一个路由方法，供客户端使用
	AddRouter(msgID uint32, router IRouter)
	//获取当前服务器的连接管理器
	GetConnManager() IConnManager
	//注册OnConnStart() 钩子函数
	SetOnConnStart(hook func(IConnection))
	//注册OnConnStop() 钩子函数
	SetOnConnStop(hook func(IConnection))
	//执行OnConnStart() 钩子函数
	CallOnConnStart(conn IConnection)
	//执行OnConnStop() 钩子函数
	CallOnConnStop(conn IConnection)
}
