package lesson5

import "DataStructure2/utils"

type SPStack struct {
	common *utils.Stack // 用于push和pop的栈
	min    *utils.Stack // 用于getMin的栈
}

func NewSPStack() *SPStack {
	return &SPStack{
		common: utils.NewStack(),
		min:    utils.NewStack(),
	}
}

func (ss *SPStack) Push(data int) *SPStack {
	ss.common.Push(data)
	minData := ss.min.Top()
	if minData == nil {
		ss.min.Push(data)
		return ss
	}

	if data < minData.(int) {
		ss.min.Push(data)
	} else {
		ss.min.Push(minData)
	}
	return ss
}

func (ss *SPStack) Pop() int {
	data := ss.common.Pop()
	ss.min.Pop()

	if data != nil {
		return data.(int)
	} else {
		return -1
	}
}

func (ss *SPStack) BackMin() int {
	min := ss.min.Top()
	if min != nil {
		return min.(int)
	} else {
		return -1
	}
}
