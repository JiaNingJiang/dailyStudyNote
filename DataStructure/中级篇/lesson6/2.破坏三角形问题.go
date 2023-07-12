package lesson6

func BrokenTriangle(n int) int {
	fibSet := make(map[int]struct{}, 0)

	var fibCount int64 = 1
	for {
		curFib := Fibonacci(fibCount)
		if curFib > n {
			break
		}

		fibCount++
		fibSet[curFib] = struct{}{}
	}
	remove := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		if _, ok := fibSet[i]; !ok {
			remove = append(remove, i)
		}
	}

	return len(remove)
}
