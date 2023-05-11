package linkList

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

func (s *Stack) Push(data interface{}) {
	s.Items = append(s.Items, data)
	s.Len++
	s.Head++
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
	return s.Items[s.Head-1]
}

func IsNotPalindrome(ll *LinkList) bool {
	s := NewStack()

	if ll.Len < 2 {
		return false
	}
	current := ll.Head
	for { // 第一次循环，将链表从头到尾的节点数据压入到栈中
		if current == nil {
			break
		}
		s.Push(current.Data)
		current = current.Next
	}
	current = ll.Head
	for { // 第二次循环，判断原链表是否是回文结构
		if current == nil { // 全部相等，则链表为回文结构
			return true
		}
		data := s.Pop()
		if NodeDataCompare(data, current.Data, TypeInt) != 0 { // 比较栈弹出的元素和链表当前节点数据是否相等
			return false // 一旦不相等，返回false
		}
		current = current.Next
	}
}
