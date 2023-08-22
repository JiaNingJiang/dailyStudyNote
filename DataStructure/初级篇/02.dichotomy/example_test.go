package dichotomy

import (
	"fmt"
	"testing"
)

func TestIsExist(t *testing.T) {
	arr := []int{0, 2, 4, 6, 8}
	target := 7
	exist, _ := IsExist(arr, target, 0, (len(arr)-1)/2, len(arr)-1)
	if exist {
		fmt.Printf("目标数字 %d 存在于数组中\n", target)
	} else {
		fmt.Printf("目标数字 %d 不存在于数组中\n", target)
	}
}

func TestAlmostLeft(t *testing.T) {
	arr := []int{0, 2, 4, 6, 8}
	target := 5

	if target < arr[0] || target >= arr[len(arr)-1] {
		t.Fatal("越界，目标值必须在 0 ~ len(arr)-1 之内")
	}

	index, value := AlmostLeft(arr, target, 0, (len(arr)-1)/2, len(arr)-1)

	fmt.Printf("大于(%d)的最左侧位置:%d 数值(%d)\n", target, index, value)
}

func TestLocalMin(t *testing.T) {
	arr1 := []int{0, 2, 4, 6, 8}
	arr2 := []int{8, 6, 4, 2, 0}
	arr3 := []int{7, 5, 2, 3, 4, 8, 9}

	index, value := LocalMin(arr1)
	fmt.Printf("arr1的一个局部最小值为:%d 下标(%d)\n", value, index)

	index, value = LocalMin(arr2)
	fmt.Printf("arr2的一个局部最小值为:%d 下标(%d)\n", value, index)

	index, value = LocalMin(arr3)
	fmt.Printf("arr3的一个局部最小值为:%d 下标(%d)\n", value, index)
}
