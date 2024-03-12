package znet

import "GO_Demo/zinx/ziface"

type Request struct {
	//与客户端建立好的连接
	conn ziface.IConnection //用的抽象类，而非结构体形式的Connection

	//客户端请求的数据,Message包(TLV格式)
	data ziface.IMessage
}

//获取当前请求包对应的连接(是哪一个客户端发起的该请求)
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

//获取请求包中包含的数据(来自于客户端的数据)
func (r *Request) GetData() []byte {
	return r.data.(*Message).Data
}

//获取请求包的消息ID
func (r *Request) GetMsgID() uint32 {
	return r.data.(*Message).ID
}
