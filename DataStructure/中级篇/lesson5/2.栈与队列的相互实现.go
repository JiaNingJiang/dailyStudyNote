package lesson5

import (
	"DataStructure2/utils"
)

type StackByQueue struct {
	MainQueue *utils.Queue
	AimQueue  *utils.Queue
}

func NewStackByQueue() *StackByQueue {
	return &StackByQueue{
		MainQueue: utils.NewQueue(),
		AimQueue:  utils.NewQueue(),
	}
}

func (sbq *StackByQueue) Push(data interface{}) *StackByQueue {
	sbq.AimQueue.Push(data)

	mainQLen := sbq.MainQueue.Len

	for i := 0; i < mainQLen; i++ {
		ele := sbq.MainQueue.Pop()
		sbq.AimQueue.Push(ele)
	}
	aimQLen := sbq.AimQueue.Len
	for i := 0; i < aimQLen; i++ {
		ele := sbq.AimQueue.Pop()
		sbq.MainQueue.Push(ele)
	}

	return sbq
}

func (sbq *StackByQueue) Pop() interface{} {
	return sbq.MainQueue.Pop()
}

type QueueByStack struct {
	PushStack *utils.Stack // 加入元素的栈
	PopStack  *utils.Stack // 输出元素的栈
}

func NewQueueByStack() *QueueByStack {
	return &QueueByStack{
		PushStack: utils.NewStack(),
		PopStack:  utils.NewStack(),
	}
}

func (qbs *QueueByStack) Push(data interface{}) *QueueByStack {
	qbs.PushStack.Push(data)
	return qbs
}

func (qbs *QueueByStack) Pop() interface{} {
	popSLen := qbs.PopStack.Len
	if popSLen != 0 {
		return qbs.PopStack.Pop()
	}
	pushSLen := qbs.PushStack.Len
	for i := 0; i < pushSLen; i++ {
		data := qbs.PushStack.Pop()
		qbs.PopStack.Push(data)
	}
	return qbs.PopStack.Pop()
}
