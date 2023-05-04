package recursion

import (
	"fmt"
	"testing"
)

func TestGetMax(t *testing.T) {
	arr := []int{3, 6, 7, 1, 7, 4, 0}

	fmt.Printf("最大值为:%d\n", GetMax(arr))
}
