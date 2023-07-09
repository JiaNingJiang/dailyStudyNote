package lesson5

import (
	"fmt"
	"testing"
)

func TestNewCoffeePro(t *testing.T) {
	coffeeConsume := []int{2, 3, 5}
	cfp := NewCoffeePro(4, coffeeConsume, 2, 3)

	for i := 0; i < cfp.Man; i++ {
		cfp.MakeCoffee()
	}

	fmt.Printf("最后结束的时间: %d\n", cfp.MinTime())
}
