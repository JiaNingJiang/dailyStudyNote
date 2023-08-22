package simpleSort

import (
	"DataStructure/03.comparator"
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{17, 13, 25, 72, 100}
	BubbleSort(arr)

	fmt.Println("冒泡排序后：", arr)

	comparator.Comparator(BubbleSort)
}
