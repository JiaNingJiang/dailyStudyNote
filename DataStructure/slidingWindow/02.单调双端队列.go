package slidingWindow

type MonotonicDeque struct {
	Monotony  bool          // true为单调递增   false为单调递减
	dataSet   []interface{} // 存储原始数组（这个数组只会增长，不会减小。也就是说只能存储加入的元素，不能将元素弹出）
	Len       int           // 每添加一个新元素时 + 1（记录下一个要添加的新元素在dataSet中的下标）
	Start     int           // 每弹出一个旧元素时 +1（记录下一个要弹出的旧元素在dataSet中的下标）
	DequeBase *Deque        // 双端队列(存储下标值)
}

func NewMonotonicDeque(monotony bool) *MonotonicDeque {
	return &MonotonicDeque{
		Monotony:  monotony,
		dataSet:   make([]interface{}, 0),
		DequeBase: NewDeque(),
	}
}

// 添加一个新元素的记录
func (mdq *MonotonicDeque) PushNew(data interface{}) {
	if mdq.Monotony { // 递增队列
		for {
			if mdq.DequeBase.Len == 0 || data.(int) <= mdq.dataSet[mdq.DequeBase.End().(int)].(int) {
				mdq.DequeBase.PushEnd(mdq.Len) // 将元素下标插入到双端队列队尾
				break
			}
			if data.(int) > mdq.dataSet[mdq.DequeBase.End().(int)].(int) { // 新加的元素大于或等于队尾元素，则将队尾元素弹出
				mdq.DequeBase.PopEnd()
			}
		}

	} else { //递减队列
		for {
			if mdq.DequeBase.Len == 0 || data.(int) >= mdq.dataSet[mdq.DequeBase.End().(int)].(int) {
				mdq.DequeBase.PushEnd(mdq.Len) // 将元素下标插入到双端队列队尾
				break
			}
			if data.(int) < mdq.dataSet[mdq.DequeBase.End().(int)].(int) { // 新加的元素大于或等于队尾元素，则将队尾元素弹出
				mdq.DequeBase.PopEnd()
			}
		}
	}
	mdq.dataSet = append(mdq.dataSet, data)
	mdq.Len++
}

// 返回当前存在记录内的最大值(递增队列)、最小值(递减队列)
func (mdq *MonotonicDeque) Pop() interface{} {
	data := mdq.DequeBase.Front()
	if data != nil {
		return mdq.dataSet[data.(int)]
	} else {
		return nil
	}

}

// 移除一个旧元素的记录
func (mdq *MonotonicDeque) RemoveOldRecord() interface{} {

	if mdq.Start == mdq.Len {
		return nil
	}

	if mdq.Start == mdq.DequeBase.Front().(int) {
		mdq.Start++
		return mdq.dataSet[mdq.DequeBase.PopFront().(int)]
	} else {
		mdq.Start++
		return nil
	}
}
