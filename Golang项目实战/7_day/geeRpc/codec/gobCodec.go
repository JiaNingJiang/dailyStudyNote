package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

type GobCodec struct {
	conn io.ReadWriteCloser // gob.Decoder的读缓存,服务端将从此处读取RPC消息头和消息体gob编码后的字节流
	buf  *bufio.Writer      // gob.Encoder的写缓冲,负责存储gob编码后的字节流
	dec  *gob.Decoder
	enc  *gob.Encoder
}

// 关闭读写器
func (gc *GobCodec) Close() error {
	return gc.conn.Close()
}

// 从GobCodec.conn中读取RPC消息的消息头的gob编码字节流,将其解码后存储到h(*Header)中
func (gc *GobCodec) ReadHeader(h *Header) error {
	return gc.dec.Decode(h)
}

// 从GobCodec.conn中读取RPC消息的消息体的gob编码字节流,将其解码后存储到body(interface{})中
func (gc *GobCodec) ReadBody(body interface{}) error {
	return gc.dec.Decode(body)
}

// 将RPC消息的消息头+消息体进行gob编码,写入到gob.Encoder指定的缓存(GobCodec.buf)中
func (c *GobCodec) Write(h *Header, body interface{}) (err error) {
	defer func() {
		_ = c.buf.Flush() // 程序结束前,将gob编码后的数据Flush到缓存中
		if err != nil {
			_ = c.Close() // 如果编码过程出现问题,则需要直接关闭读写器
		}
	}()
	if err := c.enc.Encode(h); err != nil {
		log.Println("rpc codec: gob error encoding header:", err)
		return err
	}
	if err := c.enc.Encode(body); err != nil {
		log.Println("rpc codec: gob error encoding body:", err)
		return err
	}
	return nil
}

var _ Codec = (*GobCodec)(nil) // 这一行的作用看起来像是: 如果GobCodec没有实现Codec接口,则会在这里直接报错,而不是在后续的程序中运行出错

func NewGobCodec(conn io.ReadWriteCloser) Codec {
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

// 创建使用Gob编码的Codec
func init() {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
