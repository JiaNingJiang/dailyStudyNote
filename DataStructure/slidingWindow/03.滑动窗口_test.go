package slidingWindow

import (
	"fmt"
	"testing"
)

func TestSildingWindow(t *testing.T) {
	win := NewWindow(true, 4)

	win.Push(3)
	win.Push(2)
	win.Push(4)
	win.Push(6)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 3 2 4 6

	win.Push(3)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 2 4 6 3

	win.Push(5)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 4 6 3 5

	win.Push(5)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 6 3 5 5

	win.Push(3)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 3 5 5 3

	win.Push(3)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 5 5 3 3

	win.Push(3)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 5 3 3 3

	win.Push(3)
	fmt.Printf("下标:%d~%d   最大值:%d\n", win.Start, win.End-1, win.BackPeak()) // 3 3 3 3
}
