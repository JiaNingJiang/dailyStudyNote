package lesson8

import (
	"fmt"
	"testing"
)

func TestGetPostOrder(t *testing.T) {
	pre := []int{1, 2, 4, 5, 3, 6, 7}
	in := []int{4, 2, 5, 1, 6, 3, 7}

	// 4 5 2 6 7 3 1
	fmt.Println("后序遍历序列： ", GetPostOrder(pre, in))
}
