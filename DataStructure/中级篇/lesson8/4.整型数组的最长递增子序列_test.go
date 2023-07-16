package lesson8

import (
	"fmt"
	"testing"
)

func TestLongestIncrease(t *testing.T) {
	//arr := []int{3, 1, 2, 6, 3, 4, 0}
	//arr := []int{7, 6, 5, 4, 3, 2, 1}
	arr := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println("最长递增子序列的长度为: ", LongestIncrease(arr))
}
