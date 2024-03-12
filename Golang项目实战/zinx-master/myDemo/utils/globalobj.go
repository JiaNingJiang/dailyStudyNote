package utils

import (
	"GO_Demo/zinx/ziface"
	"encoding/json"
	"io/ioutil"
)

//存放一切有关zinx框架的全局参数，供其他模块使用，一切参数可由用户通过zinx.json由用户自行进行配置

type GlobalObj struct {

	//服务器相关
	TcpServer ziface.IServer //当前Zinx全局的server对象
	Host      string         //当前服务器监听的IP地址
	TcpPort   int            //当前服务器监听的端口号
	Name      string         //服务器名称

	//Zinx相关
	ZinxVersion   string //Zinx版本号
	MaxConn       int    //最多允许的连接客户端数目
	MaxPacketSize uint32 //服务器接收客户端数据包最大字节数

	WorkerPoolSize    uint32 //读写协程之间的消息队列处理模块中worker协程的数目
	MaxWorkerQueueLen uint32 //每个worker协程的消息队列的队列长度(也就是一次性承载的消息数)
}

//定义一个可供全局使用的GlobalObj对象
var GlobalObject *GlobalObj

//读取配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, g); err != nil { //按照json格式解析，配置 GlobalObject
		panic(err)
	}
}

//导入utils包时，自动调用init函数初始化GlobalObject(从conf/zinx.json读取相关配置参数)
func init() {
	//默认参数
	GlobalObject = &GlobalObj{
		Name:              "ZinxServerApp",
		ZinxVersion:       "ZinxV0.3",
		Host:              "0.0.0.0",
		TcpPort:           8080,
		MaxConn:           1,
		MaxPacketSize:     4096,
		WorkerPoolSize:    10,
		MaxWorkerQueueLen: 1024,
	}

	//读取配置文件
	GlobalObject.Reload()
}
