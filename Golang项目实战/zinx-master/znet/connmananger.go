package znet

import (
	"GO_Demo/zinx/ziface"
	"errors"
	"fmt"
	"sync"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理的所有连接集合
	connLock    sync.RWMutex                  //保护连接集合的读写锁(因为存在多协程的写操作)
}

//创建一个ConnManager对象
func NewConnManager() ziface.IConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

//增加连接
func (cm *ConnManager) Add(conn ziface.IConnection) {

	//1.加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock() //3.解锁
	//2.在map中添加连接
	cm.connections[conn.GetConnID()] = conn

	fmt.Printf("connID :%v has added into ConnManager successfully, ConnLen :%v\n", conn.GetConnID(), cm.NumConn())
}

//删除连接
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	//1.加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock() //3.解锁
	//2.在map中删除连接
	delete(cm.connections, conn.GetConnID())

	fmt.Printf("connID :%v has remove from ConnManager successfully, ConnLen :%v\n", conn.GetConnID(), cm.NumConn())
}

//根据connID返回连接
func (cm *ConnManager) GetConn(connID uint32) (ziface.IConnection, error) {

	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	conn, ok := cm.connections[connID] //搜索是否有该连接，有则返回
	if !ok {
		fmt.Printf("connID:%v is not found in the ConnManager\n", connID)
		return nil, errors.New("conn is not found in the ConnManager")
	}
	return conn, nil
}

//得到当前连接总数
func (cm *ConnManager) NumConn() int {
	return len(cm.connections)
}

//清除并终止所有连接
func (cm *ConnManager) Clear() {

	//cm.connLock.Lock()   不需要加锁，因为conn.Stop()会调用remove()方法,这里加锁会导致死锁
	//defer cm.connLock.Unlock()

	for connID, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, connID)
	}

	fmt.Println("[server] All connnection is disconnect!")
}
