package lesson6

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	var n int64 = 20

	fmt.Printf("第%d个斐波那契数列: %d\n", n, Fibonacci(n))

	fmt.Printf("第%d个斐波那契数列: %d\n", n, fibonacci(int(n)))
}

func fibonacci(n int) int {
	sequence := []int{0, 1}

	if n <= 1 {
		if n == 0 {
			return 0
		}
		if n == 1 {
			return 1
		}
	}

	for i := 2; i <= n; i++ {
		next := sequence[i-1] + sequence[i-2]
		sequence = append(sequence, next)
	}

	return sequence[n]
}
