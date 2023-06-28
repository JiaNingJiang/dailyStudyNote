package slidingWindow

type Deque struct {
	Items []interface{}
	Len   int
}

func NewDeque() *Deque {
	return &Deque{
		Items: make([]interface{}, 0),
		Len:   0,
	}
}

// 向队尾添加元素
func (dq *Deque) PushEnd(data interface{}) {
	dq.Items = append(dq.Items, data)
	dq.Len++
}

// 向队首添加元素
func (dq *Deque) PushFront(data interface{}) {
	dq.Items = append([]interface{}{data}, dq.Items...)
	dq.Len++
}

// 从队首弹出元素
func (dq *Deque) PopFront() interface{} {
	if dq.Len == 0 {
		return nil
	}
	data := dq.Items[0]
	dq.Items = dq.Items[1:] // 删除出队列的元素
	dq.Len--
	return data
}

// 从队尾弹出元素
func (dq *Deque) PopEnd() interface{} {
	if dq.Len == 0 {
		return nil
	}
	data := dq.Items[len(dq.Items)-1]
	dq.Items = dq.Items[:len(dq.Items)-1] // 删除出队列的元素
	dq.Len--
	return data
}

// 返回队尾元素(不弹出)
func (dq *Deque) End() interface{} {
	if dq.Len == 0 {
		return nil
	}

	return dq.Items[len(dq.Items)-1]
}

// 返回队首元素(不弹出)
func (dq *Deque) Front() interface{} {
	if dq.Len == 0 {
		return nil
	}
	return dq.Items[0]
}
