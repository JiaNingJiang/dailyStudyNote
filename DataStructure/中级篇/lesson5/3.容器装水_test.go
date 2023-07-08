package lesson5

import (
	"fmt"
	"testing"
)

func TestHoldWater(t *testing.T) {
	vessel := []int{3, 1, 2, 5, 2, 4}

	fmt.Println("容器蓄水量: ", HoldWater(vessel))

	fmt.Println("容器蓄水量: ", HoldWaterByPoint(vessel))
}
