package main

import "fmt"

func findTwoSwapped(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] < nums[i] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	x, y := nums[index1], nums[index2]
	return x, y
}

func main() {
	nums := []int{1, 6, 3, 4, 5, 2, 7}

	fmt.Println(findTwoSwapped(nums))
}
