package ziface

type IDataPack interface {

	//获取message包的头部长度
	MsgHeadLen() uint32
	//封包
	Pack(IMessage) ([]byte, error)
	//解包
	UnPack([]byte) (IMessage, error)
}
