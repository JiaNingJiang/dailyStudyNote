package bitOperation

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var a, b int32 = 15, 11
	fmt.Printf("%d + %d = %d\n", a, b, Add(a, b))

	Comparator(Add, func(i1 int32, i2 int32) int32 {
		return i1 + i2
	})
}

func TestSub(t *testing.T) {
	var a, b int32 = 15, 11
	fmt.Printf("%d - %d = %d\n", a, b, Sub(b, a))

	Comparator(Sub, func(i1 int32, i2 int32) int32 {
		return i1 - i2
	})
}

func TestMul(t *testing.T) {
	var a, b int32 = -15, 11
	fmt.Printf("%d * %d = %d\n", a, b, Mul(b, a))

	Comparator(Mul, func(i1 int32, i2 int32) int32 {
		return i1 * i2
	})
}

func TestDiv(t *testing.T) {
	var a, b int32 = -16, 4
	fmt.Printf("%d / %d = %d\n", a, b, Div(a, b))

	Comparator(Div, func(i1 int32, i2 int32) int32 {
		return i1 / i2
	})
}
