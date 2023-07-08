package lesson4

import (
	"fmt"
	"testing"
)

func TestPrimefac(t *testing.T) {
	num := 17
	fmt.Printf("(%d)的质因数集合：%v\n", num, Primefac(num))
}
