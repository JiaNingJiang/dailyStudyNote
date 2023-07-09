package lesson6

import (
	"fmt"
	"math"
	"testing"
)

func selfMulInt(self interface{}, factor interface{}) interface{} {
	if self == nil {
		var selfInt int64 = 1
		factorInt := factor.(int64)
		return selfInt * factorInt
	} else {
		selfInt := self.(int64)
		factorInt := factor.(int64)
		return selfInt * factorInt
	}
}

func TestEffectivePow(t *testing.T) {
	var x int64 = 15
	var y int64 = 14

	powRes := EffectivePow(x, y, selfMulInt)
	powResStandard := int(math.Pow(float64(x), float64(y)))
	fmt.Printf("幂运算结果: %d\n", powRes)
	fmt.Printf("标准幂运算结果: %d\n", powResStandard)
}
