package greedy

import (
	"fmt"
	"testing"
)

func TestSlicingGoldBar(t *testing.T) {
	arr := []int{10, 30, 20}
	fmt.Println("切割金块所需的最小金额:", SlicingGoldBar(arr))
}
