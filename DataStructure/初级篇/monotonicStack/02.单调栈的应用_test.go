package monotonicStack

import (
	"fmt"
	"testing"
)

func TestProcess1(t *testing.T) {
	arr := []int{5, 3, 2, 1, 6, 7, 8, 4}

	fmt.Println("最大系数为:", process1(arr))
}
