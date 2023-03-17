package codec

import (
	"io"
)

// RPC消息头
type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // 客户端发送RPC消息时,每一个消息都需要携带一个序列号
	Error         string // 服务端回复的错误信息
}

// 负责对RPC消息进行编解码的接口
type Codec interface {
	io.Closer                         // 关闭读写缓存
	ReadHeader(*Header) error         // 解码,读取消息头
	ReadBody(interface{}) error       // 解码,读取消息体
	Write(*Header, interface{}) error // 编码,写入消息头和消息体
}

// Codec接口的构造函数(通过传入的实现了io.ReadWriteCloser接口的对象来构造Codec)
type NewCodecFunc func(io.ReadWriteCloser) Codec

// 用于指定Codec的编解码方法(gob/json)
type Type string

const (
	GobType  Type = "application/gob"
	JsonType Type = "application/json" // not implemented
)

// 保存已有的所有类型的Codec
var NewCodecFuncMap map[Type]NewCodecFunc
