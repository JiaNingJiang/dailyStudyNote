package lesson5

import (
	"fmt"
	"testing"
)

func TestLeftAndRightMaxDiff(t *testing.T) {
	//arr := []int{4, 5, 2, 1, 1}   //最大值在中间
	//arr := []int{5, 4, 2, 1, 1}   // 最大值在最左侧
	arr := []int{1, 4, 2, 1, 5} // 最大值在最右侧
	fmt.Println("最大差值: ", LeftAndRightMaxDiff(arr))
}
