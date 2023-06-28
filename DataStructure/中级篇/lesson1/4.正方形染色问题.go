package lesson1

import "math"

// RR……GGG
func MinPain(lattice []string) int {
	N := len(lattice)
	dryCount := math.MaxInt

	for lLattice := 1; lLattice <= N; lLattice++ { // 统计左右区域不同大小的情况下，需要染色的格子数
		if lLattice == 0 { // 左边的格子数为0
			// 统计lattice[0...N-1]一共有多少个R，将其全部染成G
			count := countRed1(lattice, 0, N-1)
			dryCount = int(math.Min(float64(dryCount), float64(count)))
		} else if lLattice == N { //左边格子数为N
			// 统计lattice[0...N-1]一共有多少个G，将其全部染成R
			count := countGreen1(lattice, 0, N-1)
			dryCount = int(math.Min(float64(dryCount), float64(count)))
		} else { // 左边的格子数为 1~N-1
			// 统计左区域lattice[0...lLattice-1]一共有多少个G，将其全部染成R + 右区域lattice[lLattice……N-1]一共有多少个R，将其全部染成G
			count := countGreen1(lattice, 0, lLattice-1) + countRed1(lattice, lLattice, N-1)
			dryCount = int(math.Min(float64(dryCount), float64(count)))
		}
	}
	return dryCount
}

// 统计lattice指定范围内红色格子的数量
func countRed1(lattice []string, left, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if lattice[i] == "R" {
			count++
		}
	}
	return count
}

// 统计lattice指定范围内绿色格子的数量
func countGreen1(lattice []string, left, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if lattice[i] == "G" {
			count++
		}
	}
	return count
}

func MinPain2(lattice []string) int {
	N := len(lattice)
	dryCount := math.MaxInt

	help1 := make([]int, N, N) // help1[i]会统计lattice[0]~lattice[i] ‘G’的数量
	help2 := make([]int, N, N) // help2[i]会统计lattice[i]~lattice[N-1] ‘R’的数量

	Gcount := 0
	for right := 0; right < N; right++ {
		if lattice[right] == "G" {
			Gcount++
		}
		help1[right] = Gcount
	}

	Rcount := 0
	for left := N - 1; left >= 0; left-- {
		if lattice[left] == "R" {
			Rcount++
		}
		help2[left] = Rcount
	}

	for lLattice := 1; lLattice <= N; lLattice++ { // 统计左右区域不同大小的情况下，需要染色的格子数
		if lLattice == 0 { // 左边的格子数为0
			// 统计lattice[0...N-1]一共有多少个R，将其全部染成G
			count := help2[0]
			dryCount = int(math.Min(float64(dryCount), float64(count)))
		} else if lLattice == N { //左边格子数为N
			// 统计lattice[0...N-1]一共有多少个G，将其全部染成R
			count := help1[N-1]
			dryCount = int(math.Min(float64(dryCount), float64(count)))
		} else { // 左边的格子数为 1~N-1
			// 统计左区域lattice[0...lLattice-1]一共有多少个G，将其全部染成R + 右区域lattice[lLattice……N-1]一共有多少个R，将其全部染成G
			count := help1[lLattice-1] + help2[lLattice]
			dryCount = int(math.Min(float64(dryCount), float64(count)))
		}
	}
	return dryCount
}
