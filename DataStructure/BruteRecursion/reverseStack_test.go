package BruteRecursion

import (
	"DataStructure/linkList"
	"fmt"
	"testing"
)

func TestReverseStack(t *testing.T) {
	stack := linkList.NewStack()

	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	fmt.Println(stack.Items)

	ReverseStack(stack)

	fmt.Println(stack.Items)
}
