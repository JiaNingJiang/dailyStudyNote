package service

import (
	"fmt"
	"reflect"
	"testing"
)

type Foo int

type Args struct{ Num1, Num2 int }

func (f Foo) Sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

// it's not a exported Method
func (f Foo) sum(args Args, reply *int) error {
	*reply = args.Num1 + args.Num2
	return nil
}

func _assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v...))
	}
}

// 测试结构体是否成功注册为rpc服务
func TestNewService(t *testing.T) {
	var foo Foo
	s := NewService(&foo)
	_assert(len(s.Method) == 1, "wrong service Method, expect 1, but got %d", len(s.Method))
	mType := s.Method["Sum"]
	_assert(mType != nil, "wrong Method, Sum shouldn't nil")
}

// 测试rpc服务能否能够调用成功
func TestMethodType_Call(t *testing.T) {
	var foo Foo
	s := NewService(&foo)
	mType := s.Method["Sum"]

	argv := mType.NewArgv()
	replyv := mType.NewReplyv()
	argv.Set(reflect.ValueOf(Args{Num1: 1, Num2: 3})) // 设置传入参数
	err := s.Call(mType, argv, replyv)                // 调用rpc服务
	_assert(err == nil && *replyv.Interface().(*int) == 4 && mType.NumCalls() == 1, "failed to call Foo.Sum")
}
