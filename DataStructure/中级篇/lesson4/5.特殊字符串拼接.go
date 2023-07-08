package lesson4

func SpStr(length int) int {
	primeSet := Primefac(length)

	res := 0
	for _, v := range primeSet {
		res += (v - 1)
	}
	return res
}

// 质因数分解
func Primefac(n int) []int {
	primeSet := make([]int, 0)

	dividend := n // 被除数
	for {
		if dividend == 1 { // 被除数变为1，完成质因数分解
			return primeSet
		}
		for div := 2; div <= dividend; div++ { // 除数从2开始，且必须小于等于被除数
			if dividend%div == 0 { //能除尽
				primeSet = append(primeSet, div)
				dividend = dividend / div
				break
			}
		}
	}
}
