package znet

import "GO_Demo/zinx/ziface"

//这是一个基类路由，在实现具体的Router是，需要继承这个BaseRouter，然后根据需要对这个BaseRouter的三个方法进行重写即可
type BaseRouter struct{}

//之所以在基类中不具体三个方法的内容，是因为每一个Router的实现要求不同，需要按照自身的要求重写这几个方法

//在处理Conn业务之前的钩子Hook方法
func (r *BaseRouter) PreHandle(ziface.IRequest) {

}

//处理Conn业务的主方法Hook
func (r *BaseRouter) Handle(ziface.IRequest) {

}

//处理完Conn业务之后的钩子Hook方法
func (r *BaseRouter) PostHandle(ziface.IRequest) {

}
