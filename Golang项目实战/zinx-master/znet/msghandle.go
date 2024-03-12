package znet

import (
	"GO_Demo/zinx/myDemo/utils"
	"GO_Demo/zinx/ziface"
	"fmt"
)

type MsgHandle struct {

	//map表，可以根据msgID索引相应的路由方法
	Apis map[uint32]ziface.IRouter
	//负责在woker协程中取任务的消息队列(消息就是服务器读协程要发送给写协程的内容)
	TaskQueue []chan ziface.IRequest
	//业务工作Worker池的worker协程的数目(一个worker协程拥有一个消息队列)
	WorkerPoolSize uint32
}

//初始化一个MsgHandle对象(map需要申请空间)
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,                               //worker协程的数目最好是可以用用户在全局配置中自行进行设置
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize), //消息队列的数目应该与worker数目相同
	}
}

//根据不同的客户端发送来的请求request调度/执行相应的router消息处理方法
func (mh *MsgHandle) DoMessageHandler(request ziface.IRequest) {

	//1.从request中提取msgID,再根据msgID获取注册的路由
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Printf("msgID :%v need register a new router\n", request.GetMsgID())
	}
	//2. 调用路由方法，将request作为参数进行传递
	handler.PostHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}

//为指定的MsgID注册相应的消息处理方法
func (mh *MsgHandle) AddRouter(msgID uint32, router ziface.IRouter) {

	//1.首先查看对应msgID是否已经注册路由方法
	if _, ok := mh.Apis[msgID]; ok {
		fmt.Println("msgID :", msgID, " need to register Corresponding methods")
		return
	}
	//2.为msgID注册对应的路由方法
	mh.Apis[msgID] = router
	fmt.Println("msgID:", msgID, " have registered methods successfully!")
}

//开启创建一个Worker的工作池(此方法需要暴露在外，在服务器启动时调用一次)
func (mh *MsgHandle) StartWorkerPool() {

	fmt.Println("[server]Start create worker pool")
	//每一个协程拥有一个属于自己的消息队列
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerQueueLen)
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}

}

//让一个worker协程阻塞等待自己负责的消息队列中的数据，然后调用DoMessageHandler对每一条数据(Request)进行处理
func (mh *MsgHandle) StartOneWorker(workerID int, MsgQueue chan ziface.IRequest) {

	fmt.Println("[server] Worker ID :", workerID, " is started")

	//阻塞等待消息队列中的Request消息，然后调用DoMessageHandler进行处理
	for {
		select {
		case request := <-MsgQueue:
			mh.DoMessageHandler(request)
		}
	}
}

//采用一种随机调度的方式，将服务器读协程产生的Request消息交给一个worker协程处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {

	//1.采用一种合理的平均的分配方式，这里是根据消息的客户端的连接ConnID进行分配
	WorkID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Printf("Add ConnID :%v ,request MsgID :%v to WorkerID: %v\n", request.GetConnection().GetConnID(),
		request.GetMsgID(), WorkID)

	//2.将request消息发送给对应的worker协程的消息队列
	mh.TaskQueue[WorkID] <- request
}
