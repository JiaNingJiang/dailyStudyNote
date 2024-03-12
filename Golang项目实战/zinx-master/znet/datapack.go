package znet

import (
	"GO_Demo/zinx/myDemo/utils"
	"GO_Demo/zinx/ziface"
	"bytes"
	"encoding/binary"
	"errors"
)

type DataPack struct{}

//实例化一个可供外部使用的对象
func NewDataPack() ziface.IDataPack {
	return &DataPack{} //接口必然都是实例对象的地址
}

//获取包的头部长度
func (dp *DataPack) MsgHeadLen() uint32 {
	//Message.ID 和 Message.Len 都是uint32 类型，总共4+4=8字节
	return 8
}

//对服务器将要发送的数据包Message进行封包
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	dataBuff := bytes.NewBuffer([]byte{}) //发送缓存区(接收器)
	//1.向包中写入消息的长度(二进制小段字节序方式写入)
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLenth()); err != nil {
		return nil, err
	}

	//2.向包中写入消息ID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	//3.向包中写入消息的内容
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

//对服务器接收到的二进制数据进行解包，解包为Message消息
func (dp *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {

	dataBuff := bytes.NewReader(binaryData) //读取器Reader

	msg := &Message{} //存放解包后的Message消息(只读取包头部字段，也就是消息长度和消息id)
	//1.读取消息长度(参数三需要是地址，因为是要将读取的数据写入参数三)
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Len); err != nil {
		return nil, err
	}
	//fmt.Println("msg.Len is :", msg.Len)
	//2.读取消息ID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.ID); err != nil {
		return nil, err
	}
	//fmt.Println("msg.ID is :", msg.ID)
	//3.判断包的消息长度是否大于规定的最大值
	if utils.GlobalObject.MaxPacketSize > 0 && msg.Len > utils.GlobalObject.MaxPacketSize {

		return nil, errors.New("too large size for message!")
	}

	return msg, nil //返回读取的消息头
}
