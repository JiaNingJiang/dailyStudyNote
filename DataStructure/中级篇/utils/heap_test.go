package utils

import (
	"fmt"
	"testing"
)

func lessInt(a, b interface{}) bool {
	aInt := a.(int)
	bInt := b.(int)

	if aInt < bInt {
		return true
	} else {
		return false
	}
}

func TestGetSortIntArrFromHeap(t *testing.T) {
	src := []interface{}{0, 1, 2, 3, 4}
	arr := NewHeap(src, true, lessInt) // 先构建大根堆
	//GetSortArrFromHeap(arr, true, lessInt) // 再完成堆排序

	fmt.Println(arr)
}

type Node struct {
	Val int
}

func lessNode(a, b interface{}) bool {
	aNode := a.(Node)
	bNode := b.(Node)

	if aNode.Val < bNode.Val {
		return true
	} else {
		return false
	}
}

func TestGetSortSpArrFromHeap(t *testing.T) {
	src := make([]interface{}, 0)
	src = append(src, Node{Val: 0})
	src = append(src, Node{Val: 1})
	src = append(src, Node{Val: 2})
	src = append(src, Node{Val: 3})
	src = append(src, Node{Val: 4})
	arr := NewHeap(src, false, lessNode) // 先构建大根堆
	//GetSortArrFromHeap(arr, false, lessNode) // 再完成堆排序

	fmt.Println(arr)
}
