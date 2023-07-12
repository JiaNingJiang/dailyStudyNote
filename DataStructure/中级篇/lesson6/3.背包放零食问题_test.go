package lesson6

import (
	"fmt"
	"testing"
)

func TestBagSnack(t *testing.T) {
	snacks := []int{1, 1, 2, 4, 5, 3, 6}

	bagSize := 4

	fmt.Println("零食放法: ", BagSnack(snacks, bagSize))

	fmt.Println("零食放法: ", BagSnackDP(snacks, bagSize))
}
