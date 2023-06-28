package simpleSort

import (
	"DataStructure/comparator"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{17, 13, 25, 72, 100}
	BubbleSort(arr)

	fmt.Println("冒泡排序后：", arr)

	comparator.Comparator(BubbleSort)
}
