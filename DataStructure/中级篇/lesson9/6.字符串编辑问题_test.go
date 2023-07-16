package lesson9

import (
	"fmt"
	"testing"
)

func TestStrEditDistance(t *testing.T) {
	str1 := "abc"
	str2 := "adc"

	icost := 5
	dcost := 3
	rcost := 100

	fmt.Println("最小编辑代价: ", StrEditDistance(str1, str2, icost, dcost, rcost))
}
