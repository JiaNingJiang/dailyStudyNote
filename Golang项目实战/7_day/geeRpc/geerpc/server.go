package geerpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"geeRpc/codec"
	"geeRpc/service"
	"io"
	"log"
	"net"
	"net/http"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	connected        = "200 Connected to Gee RPC"
	defaultRPCPath   = "/_geeprc_"
	defaultDebugPath = "/debug/geerpc"
)

const MagicNumber = 0x3bef5c

// Option消息: RPC客户端要想使用RPC服务端的服务,必须首先(而且一般只发送一次)发送Option消息与服务端进行协商,协商的内容一般会包括:
// 1.序列化方式  2.压缩方式   3.header长度   4.body长度
// 这里为了简便，我们仅为Option消息设置了序列化方式
// 在一次连接中，Option 固定在报文的最开始，Header 和 Body 可以有多个，如： | Option | Header1 | Body1 | Header2 | Body2 | ...
type Option struct {
	MagicNumber int        // MagicNumber marks this's a geerpc request
	CodecType   codec.Type // 客户端需要提前通知服务端,自身的RPC消息使用的编码类型是什么

	ConnectTimeout time.Duration // 客户端的连接超时时间（默认值为 10s）
	HandleTimeout  time.Duration // RPC请求处理超时时间 （默认值为 0，即不设限）
}

var DefaultOption = &Option{
	MagicNumber:    MagicNumber,
	CodecType:      codec.GobType, // 默认的客户端RPC编码类型是Gob
	ConnectTimeout: time.Second * 10,
}

// RPC Server类
type Server struct {
	serviceMap sync.Map // 保存当前RPC Server所有的所有service (key -> service.Name  value -> service )
}

// NewServer returns a new Server.
func NewServer() *Server {
	return &Server{}
}

// 为RPC Server 注册一个新的service
func (server *Server) Register(rcvr interface{}) error {
	s := service.NewService(rcvr)
	if _, dup := server.serviceMap.LoadOrStore(s.Name, s); dup { // 不会存储相同Name的service
		return errors.New("rpc: service already defined: " + s.Name)
	}
	return nil
}

// 专门为DefaultServer提供的服务注册函数
func Register(rcvr interface{}) error { return DefaultServer.Register(rcvr) }

// 查询获取已注册的服务(Service)以及该服务的方法(Method)
// 传入的参数 serviceMethod 是服务与服务方法的组合,如：service.Method
func (server *Server) findService(serviceMethod string) (svc *service.Service, mtype *service.MethodType, err error) {
	dot := strings.LastIndex(serviceMethod, ".")
	if dot < 0 {
		err = errors.New("rpc server: service/method request ill-formed: " + serviceMethod)
		return
	}
	serviceName, methodName := serviceMethod[:dot], serviceMethod[dot+1:] // 分别获取服务名和服务方法名
	svci, ok := server.serviceMap.Load(serviceName)                       //根据服务名获取已注册的服务Service
	if !ok {
		err = errors.New("rpc server: can't find service " + serviceName)
		return
	}
	svc = svci.(*service.Service)
	mtype = svc.Method[methodName] // 根据方法名获取已有的服务方法 Method
	if mtype == nil {
		err = errors.New("rpc server: can't find method " + methodName)
	}
	return
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
	server.serveCodec(f(conn), opt) // 使用构造函数将conn socket构建为Codec读写器，让读写器开始运行(读取RPC客户端的请求、解码、应答...)
}

// 作为对RPC客户端错误请求的应答
var invalidRequest = struct{}{}

// 负责让读写器开始运行(读取RPC客户端的请求、解码、应答...)
// 可以并发接收RPC客户端的请求,但是服务的应答必须是逐个、按序的. 因为并发容易导致多个回复报文交织在一起，客户端无法解析。在这里使用锁(sending)保证(另一种方案是在Option中对消息头,消息体的长度做出规定)
func (server *Server) serveCodec(cc codec.Codec, opt Option) {
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
		// fmt.Println("[server]  opt.HandleTimeout:", opt.HandleTimeout)
		go server.handleRequest(cc, req, sending, wg, opt.HandleTimeout)
	}
	wg.Wait()
	_ = cc.Close()
}

// request stores all information of a call
type request struct {
	h            *codec.Header // header of request
	argv, replyv reflect.Value // argv and replyv of request

	mtype *service.MethodType
	svc   *service.Service
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

// 读取客户端RPC发送的编码的字节流,返回解码后的request(包含RPC请求服务、服务方法、传入参数)
func (server *Server) readRequest(cc codec.Codec) (*request, error) {
	h, err := server.readRequestHeader(cc)
	if err != nil {
		return nil, err
	}
	req := &request{h: h}

	req.svc, req.mtype, err = server.findService(h.ServiceMethod) // 根据从RPC Client接收到的Head中包含的 ServiceMethod 获取对应的服务service和方法method
	if err != nil {
		return req, err
	}
	req.argv = req.mtype.NewArgv()
	req.replyv = req.mtype.NewReplyv()

	// make sure that argvi is a pointer, ReadBody need a pointer as parameter
	argvi := req.argv.Interface()
	if req.argv.Type().Kind() != reflect.Ptr {
		argvi = req.argv.Addr().Interface()
	}
	if err = cc.ReadBody(argvi); err != nil { // 将请求报文反序列化为第一个入参 request.argv，在这里同样需要注意 argv 可能是值类型，也可能是指针类型，所以处理方式有点差异。
		log.Println("rpc server: read body err:", err)
		return req, err
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
func (server *Server) handleRequest(cc codec.Codec, req *request, sending *sync.Mutex, wg *sync.WaitGroup, timeout time.Duration) {
	defer wg.Done()

	called := make(chan struct{}) // 用于通知当前Server端已经完成了RPC调用的处理和回复
	go func() {
		err := req.svc.Call(req.mtype, req.argv, req.replyv)
		if err != nil {
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			called <- struct{}{}
			return
		}
		server.sendResponse(cc, req.h, req.replyv.Interface(), sending)
		called <- struct{}{}
	}()

	if timeout == 0 {
		<-called
		return
	}
	select { // 要么因为RPC调用超时返回，要么因为RPC调用成功而返回
	case <-time.After(timeout):
		req.h.Error = fmt.Sprintf("rpc server: request handle timeout: expect within %s", timeout)
		server.sendResponse(cc, req.h, invalidRequest, sending)
	case <-called:
	}
}

// 将本次http通信转化为普通的tcp通信,进而实现rpc通信
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "CONNECT" { // 必须保证是CONNECT http请求
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = io.WriteString(w, "405 must CONNECT\n")
		return
	}
	conn, _, err := w.(http.Hijacker).Hijack() // 本方法的作用是：完全接管当前http连接端口，返回一个连接并关联到该socket的一个缓冲读写器(conn)。HTTP服务端将不再对连接进行任何操作，而是由conn负责接收新的数据
	if err != nil {
		log.Print("rpc hijacking ", req.RemoteAddr, ": ", err.Error())
		return
	}
	_, _ = io.WriteString(conn, "HTTP/1.0 "+connected+"\n\n") // 向rpc client 回复connect转移成功的应答
	server.ServeConn(conn)                                    // 接收来自于RPC Client的后续Option协商消息
}

// 为RPC Server设置http路由
func (server *Server) HandleHTTP() {
	http.Handle(defaultRPCPath, server)
}

func HandleHTTP() {
	DefaultServer.HandleHTTP()
}
