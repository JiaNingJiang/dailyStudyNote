package lesson1

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Equal1_5() int {
	time.Sleep(time.Nanosecond)
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn(4) + 1)
}

func TestEqualP01(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Printf(" %d ", EqualP01(Equal1_5))
		if i%10 == 0 {
			fmt.Println()
		}
	}
}

func TestEqualP1_7(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Printf(" %d ", EqualP1_7(Equal1_5))
		if i%10 == 0 {
			fmt.Println()
		}
	}
}

func TestEqualP30_59(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Printf(" %d ", EqualP30_59(Equal1_5))
		if i%10 == 0 {
			fmt.Println()
		}
	}
}
