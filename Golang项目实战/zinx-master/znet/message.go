package znet

type Message struct {
	Len  uint32 //消息长度
	ID   uint32 //消息ID
	Data []byte //消息内容
}

//初始化一个Message对象
func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		ID:   id,
		Len:  uint32(len(data)),
		Data: data,
	}
}

func (msg *Message) GetMsgID() uint32 {
	return msg.ID
}

func (msg *Message) GetMsgLenth() uint32 {
	return msg.Len
}

func (msg *Message) GetMsgData() []byte {
	return msg.Data
}
func (msg *Message) SetMsgID(id uint32) {
	msg.ID = id
}
func (msg *Message) SetMsgData(data []byte) {
	msg.Data = data
}
func (msg *Message) SetMsgLenth(len uint32) {
	msg.Len = len
}
