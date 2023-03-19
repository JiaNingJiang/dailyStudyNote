package service

import (
	"go/ast"
	"log"
	"reflect"
	"sync/atomic"
)

// 映射一个rpc服务对象的方法
type MethodType struct {
	Method    reflect.Method // rpc方法本身
	ArgType   reflect.Type   // rpc方法的请求参数
	ReplyType reflect.Type   // rpc方法的响应参数
	numCalls  uint64         // 后续统计方法调用次数时会用到
}

// 对numCalls的读取,必须是原子操作,不允许并发
func (m *MethodType) NumCalls() uint64 {
	return atomic.LoadUint64(&m.numCalls)
}

// 根据reflect.Type类型的ArgType，创建并返回与其类型对应的reflect.Value
func (m *MethodType) NewArgv() reflect.Value {
	var argv reflect.Value
	// arg may be a pointer type, or a value type
	if m.ArgType.Kind() == reflect.Ptr { // 如果使用的reflect.Type是指针,那么需要提前将其转换为值类型(.Elem())
		argv = reflect.New(m.ArgType.Elem()) //最终返回的是指针类型的实例
	} else {
		argv = reflect.New(m.ArgType).Elem() //最终返回的是值类型的实例
	}
	return argv
}

func (m *MethodType) NewReplyv() reflect.Value {
	// reply must be a pointer type
	replyv := reflect.New(m.ReplyType.Elem())
	switch m.ReplyType.Elem().Kind() { // 需要根据返回值的类型(map or slice)采取不同的措施
	case reflect.Map:
		replyv.Elem().Set(reflect.MakeMap(m.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(m.ReplyType.Elem(), 0, 0))
	}
	return replyv
}

// 映射一个rpc服务对象
type Service struct {
	Name   string                 // name 即映射的结构体的名称，比如 T，比如 WaitGroup
	Typ    reflect.Type           // typ 是结构体的类型
	Rcvr   reflect.Value          // rcvr 即结构体的实例本身，保留 rcvr 是因为在调用时需要 rcvr 作为第 0 个参数
	Method map[string]*MethodType // 存储映射的结构体的所有符合条件的方法
}

// 将一个结构体映射为一个rpc服务对象 service
func NewService(rcvr interface{}) *Service {
	s := new(Service)
	s.Rcvr = reflect.ValueOf(rcvr)
	s.Name = reflect.Indirect(s.Rcvr).Type().Name() // 获取rcvr变量的类型名(结构体名称)
	s.Typ = reflect.TypeOf(rcvr)
	if !ast.IsExported(s.Name) { // 必须确保结构体是可导出的(结构体名的首字母需要大写)
		log.Fatalf("rpc server: %s is not a valid service name", s.Name)
	}
	s.registerMethods()
	return s
}

// 将结构体的方法全部注册为rpc服务 methodType
func (s *Service) registerMethods() {
	s.Method = make(map[string]*MethodType)
	for i := 0; i < s.Typ.NumMethod(); i++ {
		method := s.Typ.Method(i)
		mType := method.Type
		if mType.NumIn() != 3 || mType.NumOut() != 1 { // 如果方法的传入参数个数不等于3,返回值个数不等于1 (第0个传入参数是结构体本身，返回值只有一个就是error类型)
			continue
		}
		if mType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() { // 返回值类型必须为error
			continue
		}
		argType, replyType := mType.In(1), mType.In(2) // 分别获取rpc方法的请求参数和响应参数
		if !isExportedOrBuiltinType(argType) || !isExportedOrBuiltinType(replyType) {
			continue
		}
		s.Method[method.Name] = &MethodType{
			Method:    method,
			ArgType:   argType,
			ReplyType: replyType,
		}
		log.Printf("rpc server: register %s.%s\n", s.Name, method.Name)
	}
}

// 检查传入参数是否是可导出类型
func isExportedOrBuiltinType(t reflect.Type) bool {
	return ast.IsExported(t.Name()) || t.PkgPath() == "" // PkgPath是非导出字段的包路径，对导出字段该字段为""。
}

// 调用rpc方法并返回error
func (s *Service) Call(m *MethodType, argv, replyv reflect.Value) error {
	atomic.AddUint64(&m.numCalls, 1) // rpc方法调用次数 +1
	f := m.Method.Func
	returnValues := f.Call([]reflect.Value{s.Rcvr, argv, replyv}) // 调用rpc方法(第一个参数s.rcvr就是结构体本身)
	if errInter := returnValues[0].Interface(); errInter != nil { // 检查是否发生err
		return errInter.(error)
	}
	return nil
}
