package geerpc

import (
	"encoding/json"
	"fmt"
	"geeRpc/codec"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

const MagicNumber = 0x3bef5c

// Option消息: RPC客户端要想使用RPC服务端的服务,必须首先(而且一般只发送一次)发送Option消息与服务端进行协商,协商的内容一般会包括:
// 1.序列化方式  2.压缩方式   3.header长度   4.body长度
// 这里为了简便，我们仅为Option消息设置了序列化方式
// 在一次连接中，Option 固定在报文的最开始，Header 和 Body 可以有多个，如： | Option | Header1 | Body1 | Header2 | Body2 | ...
type Option struct {
	MagicNumber int        // MagicNumber marks this's a geerpc request
	CodecType   codec.Type // 客户端需要提前通知服务端,自身的RPC消息使用的编码类型是什么
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType, // 默认的客户端RPC编码类型是Gob
}

// RPC Server类,暂时没有任何的成员字段
type Server struct{}

// NewServer returns a new Server.
func NewServer() *Server {
	return &Server{}
}

// 在已有的socket上提供RPC服务：接收RPC客户端的连接请求,为其提供RPC服务
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept error:", err)
			return
		}
		go server.ServeConn(conn)
	}
}

// DefaultServer 是一个默认的RPC Server实例
var DefaultServer = NewServer()

// Accept也是一个默认的RPC Server的服务启动函数,会调用 DefaultServer 为用户提供RPC服务
func Accept(lis net.Listener) { DefaultServer.Accept(lis) }

// RPC服务端提供RPC服务的核心代码.根据RPC客户端Option协商消息指定的编码类型,构造对应的Codec读写器,再让Codec读写器进行工作
func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	// 第一步:接收RPC客户端的Option协商消息
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server: options error: ", err)
		return
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	// 第二步,根据Option中指定的RPC客户端编码类型,获取对应的Codec编解码读写器的构造函数
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	server.serveCodec(f(conn)) // 使用构造函数将conn socket构建为Codec读写器，让读写器开始运行(读取RPC客户端的请求、解码、应答...)
}

// 作为对RPC客户端错误请求的应答
var invalidRequest = struct{}{}

// 负责让读写器开始运行(读取RPC客户端的请求、解码、应答...)
// 可以并发接收RPC客户端的请求,但是服务的应答必须是逐个、按序的. 因为并发容易导致多个回复报文交织在一起，客户端无法解析。在这里使用锁(sending)保证(另一种方案是在Option中对消息头,消息体的长度做出规定)
func (server *Server) serveCodec(cc codec.Codec) {
	sending := new(sync.Mutex) // make sure to send a complete response
	wg := new(sync.WaitGroup)  // wait until all request are handled
	for {
		req, err := server.readRequest(cc) // 循环读取RPC客户端的请求，获取解码后的客户端RPC请求 request(Head + Body)
		if err != nil {
			if req == nil {
				break // it's not possible to recover, so close the connection
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()
}

// request stores all information of a call
type request struct {
	h            *codec.Header // header of request
	argv, replyv reflect.Value // argv and replyv of request
}

// 读取客户端RPC发送的编码的字节流,返回解码后的RPC Head
func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

// 读取客户端RPC发送的编码的字节流,返回解码后的request(包含Head 和 Body)
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{h: h}
	// TODO: now we don't know the type of request argv
	// day 1, just suppose it's string
	req.argv = reflect.New(reflect.TypeOf(""))
	if err = cc.ReadBody(req.argv.Interface()); err != nil {
		log.Println("rpc server: read argv err:", err)
	}
	return req, nil
}

// 将客户端的RPC请求的Head 和 生成的RPC应答编码之后一起回复给客户端
func (server *Server) sendResponse(cc codec.Codec, h *codec.Header, body interface{}, sending *sync.Mutex) {
	sending.Lock()
	defer sending.Unlock()
	if err := cc.Write(h, body); err != nil {
		log.Println("rpc server: write response error:", err)
	}
}

// 处理来自于客户端的RPC请求(已完成解码的明文),生成RPC应答回复给客户端
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup) {
	// TODO, should call registered rpc methods to get the right replyv
	// day 1, just print argv and send a hello message
	defer wg.Done()
	log.Println(req.h, req.argv.Elem())                                    // 打印客户端RPC请求的消息头和消息体
	req.replyv = reflect.ValueOf(fmt.Sprintf("geerpc resp %d", req.h.Seq)) // 构造应答
	server.sendResponse(cc, req.h, req.replyv.Interface(), sending)        // 向客户端发送RPC应答
}
