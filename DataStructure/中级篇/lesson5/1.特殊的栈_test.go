package lesson5

import (
	"fmt"
	"testing"
)

func TestNewSPStack(t *testing.T) {
	ss := NewSPStack()

	ss.Push(4).Push(4).Push(2).Push(3).Push(2).Push(2)

	ss.Pop()
	fmt.Println("min : ", ss.BackMin())

	ss.Pop()
	fmt.Println("min : ", ss.BackMin())

	ss.Pop()
	fmt.Println("min : ", ss.BackMin())

	ss.Pop()
	fmt.Println("min : ", ss.BackMin())

	ss.Pop()
	fmt.Println("min : ", ss.BackMin())

	ss.Pop()
	fmt.Println("min : ", ss.BackMin())
}
