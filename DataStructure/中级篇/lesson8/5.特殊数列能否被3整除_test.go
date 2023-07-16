package lesson8

import (
	"fmt"
	"testing"
)

func TestMagicSeries(t *testing.T) {
	res := magicSeries(2, 5)

	fmt.Println(res)
}

func TestIsNotDivBy3(t *testing.T) {
	fmt.Println(isNotDivBy3(3))
	fmt.Println(isNotDivBy3(10))
	fmt.Println(isNotDivBy3(12))
	fmt.Println(isNotDivBy3(123))
}

func TestMagicSeriesDiv3(t *testing.T) {
	fmt.Println(MagicSeriesDiv3(2, 5))
}
