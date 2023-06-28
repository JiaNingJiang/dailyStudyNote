package monotonicStack

import (
	"fmt"
	"testing"
)

func TestMonotonicStackBig(t *testing.T) {
	arr := []int{5, 4, 3, 6, 1, 2, 0, 7}

	mStack := NewMonotonicStack(arr, true)

	res := mStack.Order()

	fmt.Println(res)
}

func TestMonotonicStackSmall(t *testing.T) {
	arr := []int{5, 4, 3, 6, 1, 2, 0, 7}

	mStack := NewMonotonicStack(arr, false)

	res := mStack.Order()

	fmt.Println(res)
}

func TestEqualMonotonicStackBig(t *testing.T) {
	arr := []int{5, 4, 3, 4, 5, 3, 5, 6}

	mStack := NewMonotonicStack(arr, true)

	res := mStack.Order2()

	fmt.Println(res)
}
