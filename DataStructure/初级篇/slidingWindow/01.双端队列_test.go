package slidingWindow

import (
	"fmt"
	"testing"
)

func TestDeque1(t *testing.T) {
	dq := NewDeque()

	dq.PushEnd(1)
	dq.PushEnd(2)
	dq.PushEnd(3)
	dq.PushEnd(4)
	fmt.Println(dq.Items)

	var str string
	str += fmt.Sprintf("%d", dq.PopFront())
	str += fmt.Sprintf("%d", dq.PopFront())
	str += fmt.Sprintf("%d", dq.PopFront())
	str += fmt.Sprintf("%d", dq.PopFront())
	fmt.Println(str)
	fmt.Println(dq.Items)
}

func TestDeque2(t *testing.T) {
	dq := NewDeque()

	dq.PushEnd(1)
	dq.PushEnd(2)
	dq.PushEnd(3)
	dq.PushEnd(4)
	fmt.Println(dq.Items)

	var str string
	str += fmt.Sprintf("%d", dq.PopEnd())
	str += fmt.Sprintf("%d", dq.PopEnd())
	str += fmt.Sprintf("%d", dq.PopEnd())
	str += fmt.Sprintf("%d", dq.PopEnd())
	fmt.Println(str)
	fmt.Println(dq.Items)
}

func TestDeque3(t *testing.T) {
	dq := NewDeque()

	dq.PushFront(1)
	dq.PushFront(2)
	dq.PushFront(3)
	dq.PushFront(4)
	fmt.Println(dq.Items)

	var str string
	str += fmt.Sprintf("%d", dq.PopEnd())
	str += fmt.Sprintf("%d", dq.PopEnd())
	str += fmt.Sprintf("%d", dq.PopEnd())
	str += fmt.Sprintf("%d", dq.PopEnd())
	fmt.Println(str)
	fmt.Println(dq.Items)
}

func TestDeque4(t *testing.T) {
	dq := NewDeque()

	dq.PushFront(1)
	dq.PushFront(2)
	dq.PushFront(3)
	dq.PushFront(4)
	fmt.Println(dq.Items)

	var str string
	str += fmt.Sprintf("%d", dq.PopFront())
	str += fmt.Sprintf("%d", dq.PopFront())
	str += fmt.Sprintf("%d", dq.PopFront())
	str += fmt.Sprintf("%d", dq.PopFront())
	fmt.Println(str)
	fmt.Println(dq.Items)
}
