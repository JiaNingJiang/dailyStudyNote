package lesson5

import (
	"fmt"
	"testing"
)

func TestNewStackByQueue(t *testing.T) {
	sbq := NewStackByQueue()

	sbq.Push(1).Push(2).Push(3)
	fmt.Println(sbq.Pop())
	fmt.Println(sbq.Pop())
	fmt.Println(sbq.Pop())

	fmt.Println()

	qbs := NewQueueByStack()
	qbs.Push(1).Push(2).Push(3)
	fmt.Println(qbs.Pop())
	fmt.Println(qbs.Pop())
	fmt.Println(qbs.Pop())
}
