package znet

import (
	"GO_Demo/zinx/myDemo/utils"
	"GO_Demo/zinx/ziface"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
)

type Connection struct {

	//当前的Connection隶属于哪一个server
	server ziface.IServer

	//当前TCP连接的套接字socket
	Conn *net.TCPConn

	//连接的ID
	ConnID uint32

	//当前的连接状态
	isClosed bool

	//当前链接所绑定的处理业务方法API(如何处理从客户端接收到的数据)  v0.1版本
	//handleAPI ziface.HandleFunc

	//告知当前链接已退出/停止的channel
	ExitChan chan bool

	//负责读写协程之间数据传输的管道
	MsgChan chan []byte

	//MsgHandle处理模块，根据消息id调用不同的路由API
	MsgHandler ziface.IMsgHandle

	//为客户端设计一些连接属性(比如人物的名称、HP、MP等)
	property map[string]interface{}

	//为连接属性添加的保护锁
	propertyMutex sync.RWMutex
}

//初始化一个连接模块
func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {

	c := &Connection{
		server:     server,
		Conn:       conn,
		ConnID:     connID,
		isClosed:   false,
		MsgHandler: msgHandler,
		ExitChan:   make(chan bool, 1),
		MsgChan:    make(chan []byte),
		property:   make(map[string]interface{}),
		//handleAPI: callback_api,
	}
	//将新创建的连接添加到ConnManager中
	server.GetConnManager().Add(c)

	return c
}

//读协程运行，负责循环等待接收来自相应connection连接的客户端的消息，同时负责对消息进行解包、处理，
//最后将处理后的消息通过管道发送给写协程
func (c *Connection) StartReader() {
	fmt.Println("Reader Groutine is working..")
	defer fmt.Printf("connID = %v ,Reader is exit,remote addr is %v\n", c.ConnID, c.RemoteAddr().String())
	defer c.Stop()

	for {
		//1.从客户端中读取消息头
		dp := NewDataPack()                       //获取解包对象(解包工具)
		headData := make([]byte, dp.MsgHeadLen()) //消息头缓冲区，存储读取的二进制数据

		if count, err := io.ReadFull(c.GetTCPConn(), headData); err != nil { //从套接字读取消息头
			if count == 0 || err == io.EOF { //表示客户端已将连接socket关闭(两种情况分别对应客户端异常和正常关闭socket)
				break
			}
			fmt.Println("[server] Read TLV head data is err:", err)
			continue
		}
		msg, err := dp.UnPack(headData) //对二进制数据进行解包，获得Message消息--msg
		if err != nil {
			fmt.Println("[server] Unpack head data is err:", err)
			continue
		}
		if msg.GetMsgLenth() > 0 { //message消息长度大于0，需要进行二次读取
			bodyData := make([]byte, msg.GetMsgLenth())                          //消息体缓冲区，大小等于消息头指定的长度
			if count, err := io.ReadFull(c.GetTCPConn(), bodyData); err != nil { //从套接字读取消息实体
				if count == 0 || err == io.EOF { //表示客户端已将连接socket关闭(两种情况分别对应客户端异常和正常关闭socket)
					break
				}
				fmt.Println("[server] Read TLV body data is err:", err)
				continue
			}
			msg.SetMsgData(bodyData) //将消息实体部分存入message中，打包成Request请求包，交给服务器的路由方法
		}

		//获取当前conn的Request请求包
		req := Request{
			conn: c,
			data: msg,
		}

		//将request请求包发送给Worker工作池(工作池worker数目不为0的情况下)
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MsgHandler.SendMsgToTaskQueue(&req)
		} else {
			//为当前的客户端request调用msgHandler进行路由处理
			go c.MsgHandler.DoMessageHandler(&req)
		}
	}
}

//写协程运行的函数，负责从管道中读取读写程发送来的数据，然后将其发送给客户端
func (c *Connection) StartWriter() {

	fmt.Println("Writer Groutine is working...")
	defer fmt.Printf("connID = %v ,Writer is exit,remote addr is %v\n", c.ConnID, c.RemoteAddr().String())

	//阻塞等待channel中的数据
	for {
		select {
		case data := <-c.MsgChan: //读取读协程从管道发来的数据，将其发送给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Server send data to ", c.RemoteAddr().String(), " is err :", err)
				return
			}
		case <-c.ExitChan: //服务器与客户端的Connection被Stop方法关闭了(当管道被关闭时会在写入一个默认值供读取)
			//fmt.Printf("[Writer Groutine is closed by signal : %v]\n", signal)
			return
		}
	}

}

func (c *Connection) Start() {
	fmt.Println("Conn is start ,ConnID :", c.ConnID)
	//启动当前连接的读数据业务
	go c.StartReader()
	//启动当前连接的写数据业务
	go c.StartWriter()

	//执行连接后初次需要执行的Hook函数
	c.server.CallOnConnStart(c)
}

//关闭当前TCP连接
func (c *Connection) Stop() {
	fmt.Printf("Connection is closed ,ConnID:%v\n", c.ConnID)
	if c.isClosed == true {
		return
	}
	//在tcpSocket关闭之前，执行关闭前的Hook函数
	c.server.CallOnConnStop(c)

	c.Conn.Close() //关闭连接
	c.isClosed = true

	c.server.GetConnManager().Remove(c)

	//c.ExitChan<-true
	close(c.ExitChan) //关闭管道(资源回收)
	close(c.MsgChan)  //关闭读写协程通信管道(资源回收)
}

//获取当前TCP连接的socket
func (c *Connection) GetTCPConn() *net.TCPConn {
	return c.Conn
}

//获取当前TCP连接的ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取当前连接的客户端地址
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//将封包之后的数据(TLV格式)通过管道发送给写协程
func (c *Connection) Send(msgID uint32, msgData []byte) error {
	if c.isClosed == true {
		return errors.New("connection is closed!")
	}

	msg := NewMessage(msgID, msgData) //按照TLV格式构造一条Message消息(Len字段将自动计算)
	dp := NewDataPack()               //封包对象
	binaryData, err := dp.Pack(msg)   //进行封包，得到二进制数据
	if err != nil {
		return err
	}
	//将要发送给客户端的数据通过管道发送给写协程
	c.MsgChan <- binaryData
	return nil
}

//添加连接属性
func (c *Connection) AddProperty(key string, property interface{}) {
	c.propertyMutex.Lock()
	defer c.propertyMutex.Unlock()

	c.property[key] = property
}

//删除连接属性
func (c *Connection) DeleteProperty(key string) {
	c.propertyMutex.Lock()
	defer c.propertyMutex.Unlock()

	delete(c.property, key)
}

//获取连接属性
func (c *Connection) GetProperty(key string) (interface{}, error) {
	c.propertyMutex.RLock()
	defer c.propertyMutex.RUnlock()
	if property, ok := c.property[key]; ok {
		return property, nil
	}
	return nil, errors.New("[server] Corresponding Property is not exist!")
}
