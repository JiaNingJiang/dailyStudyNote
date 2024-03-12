package ziface

/*
	连接管理模块抽象层
*/

type IConnManager interface {

	//增加连接
	Add(conn IConnection)
	//删除连接
	Remove(conn IConnection)
	//根据connID返回连接
	GetConn(connID uint32) (IConnection, error)
	//得到当前连接总数
	NumConn() int
	//清除并终止所有连接
	Clear()
}
