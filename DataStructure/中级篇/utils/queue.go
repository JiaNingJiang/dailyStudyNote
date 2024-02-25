package utils

type Queue struct {
	Items []interface{}
	Len   int
}

func NewQueue() *Queue {
	return &Queue{
		Items: make([]interface{}, 0),
		Len:   0,
	}
}

func (q *Queue) Push(data interface{}) {
	q.Items = append(q.Items, data)
	q.Len++
}

func (q *Queue) Pop() interface{} {
	data := q.Items[0]
	q.Items = q.Items[1:] // 删除出队列的元素
	q.Len--
	return data
}

func (q *Queue) Size() int {
	return q.Len
}

func (q *Queue) Empty() bool {
	return q.Size() == 0
}
