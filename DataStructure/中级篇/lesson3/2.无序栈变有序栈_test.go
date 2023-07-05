package lesson3

import (
	"DataStructure2/utils"
	"fmt"
	"testing"
)

func TestOrderization(t *testing.T) {
	stack := utils.NewStack()
	stack.Push(5).Push(2).Push(7).Push(0).Push(9)

	fmt.Printf("单调化前: ")
	for {
		if stack.Len == 0 {
			break
		}
		fmt.Printf(" %d ", stack.Pop().(int))
	}
	fmt.Println()

	fmt.Printf("单调化后: ")
	stack.Push(5).Push(2).Push(7).Push(0).Push(9)
	Orderization(stack, true)
	for {
		if stack.Len == 0 {
			break
		}
		fmt.Printf(" %d ", stack.Pop().(int))
	}
	fmt.Println()
}
