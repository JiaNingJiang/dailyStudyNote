package ziface

//处理客户端消息的抽象类
type IMessage interface {

	//获取消息ID
	GetMsgID() uint32
	//获取消息的长度
	GetMsgLenth() uint32
	//获取消息的内容
	GetMsgData() []byte
	//设置消息的ID
	SetMsgID(uint32)
	//设置消息的内容
	SetMsgData([]byte)
	//设置消息的长度
	SetMsgLenth(uint32)
}
