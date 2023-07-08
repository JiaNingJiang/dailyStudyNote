package utils

type Stack struct {
	Head  int           // 栈顶指针
	Len   int           // 当前栈大小
	Items []interface{} // 实现栈的底层切片
}

func NewStack() *Stack {
	return &Stack{
		Head:  0,
		Len:   0,
		Items: make([]interface{}, 0),
	}
}

func (s *Stack) Push(data interface{}) *Stack {
	s.Items = append(s.Items, data)
	s.Len++
	s.Head++

	return s
}

func (s *Stack) Pop() interface{} {
	if s.Len == 0 {
		return nil
	}
	s.Len--
	s.Head--
	data := s.Items[s.Head]
	s.Items = s.Items[0:s.Head] // 重要：底层数组删除这个被弹出的元素
	return data
}

func (s *Stack) Top() interface{} {
	if s.Head == 0 {
		return nil
	} else {
		return s.Items[s.Head-1]
	}
}
