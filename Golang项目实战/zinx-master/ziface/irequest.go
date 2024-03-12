package ziface

type IRequest interface {
	//获取当前请求包对于的连接(是哪一个客户端发起的该请求)
	GetConnection() IConnection

	//获取请求包中Message包中包含的数据
	GetData() []byte

	//获取请求包中Message包的消息ID
	GetMsgID() uint32
}
