package znet

import (
	"GO_Demo/zinx/myDemo/utils"
	"GO_Demo/zinx/ziface"
	"errors"
	"fmt"
	"net"
	"time"
)

//实现IServer接口
type Server struct {

	//服务器名称
	Name string
	//IP版本号
	IPVersion string
	//服务器监听的IP地址
	IP string
	//服务器监听的端口号
	Port int
	//MsgHandle处理模块，根据消息id调用不同的路由API
	MsgHandler ziface.IMsgHandle
	//连接管理器
	ConnManager ziface.IConnManager
	//服务器与客户端的连接建立之后调用的Hook函数
	OnConnStart func(conn ziface.IConnection)
	//服务器与客户端的连接关闭之后调用的Hook函数
	OnConnStop func(conn ziface.IConnection)
}

//自定义的服务器业务处理函数(与客户端连接后将调用的处理api)
func CallBackToClient(conn *net.TCPConn, data []byte, count int) error {
	//简单的将接收数据回显
	fmt.Println("[Connection Handle] CallBackToClient...")
	if _, err := conn.Write(data[:count]); err != nil {
		fmt.Println("Send buf is err:", err)
		return errors.New("CallBackToClient is error")
	}
	return nil
}

func (s *Server) Start() {

	go func() {
		fmt.Printf("[Zinx]Server Name :%s ,Server listener at IP:%s , Port:%d , is starting\n",
			utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
		fmt.Printf("[Zinx]MaxConn :%d ,MaxPacketSize:%d \n", utils.GlobalObject.MaxConn,
			utils.GlobalObject.MaxPacketSize)

		//0.开启Worker工作池
		s.MsgHandler.StartWorkerPool()

		//1.根据初始化s的字段创建一个socket(TCP Addr)
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Printf("resolve tcp addr error:%v\n", err)
			return
		}
		//2.服务器进行监听
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Printf("server listen is err:%v\n", err)
			return
		}
		var cid uint32 = 0 //每一个客户端对应的与服务器的ConnID
		//3.阻塞等待客户端连接
		for {

			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("server accept is err:", err)
				continue
			}

			//连接建立后，首先需要先判断服务器连接数是否已经达到上限
			if s.ConnManager.NumConn() >= utils.GlobalObject.MaxConn {
				//TODO:给客户端回复一个连接数已达上限的错误提示
				fmt.Println("The number of connection is up to maximum!!!")
				conn.Close()
				time.Sleep(time.Second)
				continue //等待连接数降低后再连接
			}

			//利用产生的conn连接套接字，创建一个Connection模块
			//dealConn := NewConnection(conn, cid, CallBackToClient)  v0.1版本采用自定义 CallBackToClient处理来自客户端的消息
			dealConn := NewConnection(s, conn, cid, s.MsgHandler) //v0.2版本采用Router来处理来自客户端的数据
			cid++
			//启动连接模块
			dealConn.Start()
		}
	}()

}

func (s *Server) Stop() {
	fmt.Println("[server] Server is stop")
	s.ConnManager.Clear() //回收所有连接资源
}

func (s *Server) Run() {
	s.Start() //调用Start方法，因此Start方法必须整体在一个协程中进行，否则Run方法会阻塞在Start方法中
	defer s.Stop()
	//TODO:实现一些其它的功能

	select {} //阻塞

}

//用户需自行实现一个继承与BaseRouter的新router(按需要重写三个Hook方法)作为形参传入
//接着在AddRouter()方法中为指定消息ID的Message注册该路由方法
func (s *Server) AddRouter(msgID uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgID, router)
}

//获取当前服务器的连接管理器
func (s *Server) GetConnManager() ziface.IConnManager {
	return s.ConnManager
}

//初始化创建一个server(返回一个抽象的接口)
func NewServer(name string) ziface.IServer {

	return &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMsgHandle(),
		ConnManager: NewConnManager(),
	}
}

//注册OnConnStart() 钩子函数
func (s *Server) SetOnConnStart(hook func(ziface.IConnection)) {
	s.OnConnStart = hook
}

//注册OnConnStop() 钩子函数
func (s *Server) SetOnConnStop(hook func(ziface.IConnection)) {
	s.OnConnStop = hook
}

//执行OnConnStart() 钩子函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("[server] Call On ConnStart Func!!")
		s.OnConnStart(conn)
	}

}

//执行OnConnStop() 钩子函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("[server] Call On ConnStop Func!!")
		s.OnConnStop(conn)
	}
}
