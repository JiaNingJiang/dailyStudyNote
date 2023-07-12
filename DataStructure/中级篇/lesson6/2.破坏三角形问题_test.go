package lesson6

import (
	"fmt"
	"testing"
)

func TestBrokenTriangle(t *testing.T) {
	n := 17

	fmt.Printf("至少需要扔掉 %d 根木棍，才能保证不组成三角形\n", BrokenTriangle(n))
}
