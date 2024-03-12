package ziface

type IMsgHandle interface {

	//根据不同的客户端发送来的请求request调度/执行相应的router消息处理方法
	DoMessageHandler(request IRequest)

	//为指定的MsgID注册相应的消息处理方法
	AddRouter(msgID uint32, router IRouter)

	//开启worker工作池
	StartWorkerPool()
	//采用一种随机调度的方式，将服务器读协程产生的Request消息交给一个worker协程处理
	SendMsgToTaskQueue(request IRequest)
}
