package ziface

type IRouter interface {

	//在处理Conn业务之前的钩子Hook方法
	PreHandle(IRequest)

	//处理Conn业务的主方法Hook
	Handle(IRequest)

	//处理完Conn业务之后的钩子Hook方法
	PostHandle(IRequest)
}
