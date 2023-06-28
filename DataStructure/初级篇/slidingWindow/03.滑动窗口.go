package slidingWindow

type Window struct {
	size   int             // 窗口大小
	Start  int             // 起始指针
	End    int             // 终止指针（总是指向下一个即将进队列但没进的元素）
	mDeque *MonotonicDeque // 底层单调双端队列
}

// monotony == true则表示底层是一个单调递增双端队列，每次可以从窗口获取一个最大值；否则为单调递减双端队列，每次可以获取一个最小值
func NewWindow(monotony bool, size int) *Window {
	win := new(Window)
	win.size = size
	win.mDeque = NewMonotonicDeque(monotony)

	return win
}

func (w *Window) Push(data interface{}) {
	if w.End-w.Start == w.size { // 窗口超过了上限，需要删除一个过期元素的记录
		w.mDeque.RemoveOldRecord()
		w.Start++
	}
	w.mDeque.PushNew(data)
	w.End++
}

func (w *Window) BackPeak() interface{} {
	return w.mDeque.BackPeak()
}
