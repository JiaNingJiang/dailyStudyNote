package ziface

import "net"

//定义连接模块的抽象层
type IConnection interface {

	//启动链接，让当前链接准备开始工作
	Start()

	//停止链接，结束当前链接的工作
	Stop()

	//获取当前链接的TCPAddr(连接套接字)
	GetTCPConn() *net.TCPConn

	//获取当前链接模块的链接ID
	GetConnID() uint32

	//获取远程客户端的TCP状态 (协议类型 IP:PORT)
	RemoteAddr() net.Addr

	//将封包之后的数据(TLV格式)发送给客户端
	Send(msgID uint32, msgData []byte) error

	//添加连接属性
	AddProperty(key string, property interface{})

	//删除连接属性
	DeleteProperty(key string)

	//获取连接属性
	GetProperty(key string) (interface{}, error)
}

//定义一个处理业务的API函数
type HandleFunc func(*net.TCPConn, []byte, int) error
