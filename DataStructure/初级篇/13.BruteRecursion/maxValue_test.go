package BruteRecursion

import (
	"fmt"
	"testing"
)

func TestMaxValue(t *testing.T) {
	commodities := []Commodity{
		{2, 4},
		{1, 3},
		{3, 5},
		{6, 9},
		{2, 3},
	}
	fmt.Println("f1--最大利润:", MaxValue(commodities, 10))
	fmt.Println("f2--最大利润:", MaxValue2(commodities, 10))
}
