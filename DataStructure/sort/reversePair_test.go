package sort

import (
	"fmt"
	"testing"
)

func TestReversePair(t *testing.T) {
	arr := []int{7, 6, 5, 4, 3, 2, 1}

	ReversePair(arr)

	fmt.Println(arr)
}
