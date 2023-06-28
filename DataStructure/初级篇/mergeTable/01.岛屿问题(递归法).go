package mergeTable

func CountIsland(m [][]int) int {
	if m == nil || m[0] == nil {
		return 0
	}
	line := len(m)   // 矩阵行数
	col := len(m[0]) // 矩阵列数
	island := 0      // 记录岛屿数

	for i := 0; i < line; i++ {
		for j := 0; j < col; j++ {
			if m[i][j] == 1 {
				island++
				infect(m, line, col, i, j)
			}
		}
	}
	return island
}

func infect(m [][]int, line, col int, i, j int) {
	if i < 0 || i >= line || j < 0 || j >= col || m[i][j] != 1 { // 1.访问矩阵不能越界  2.如果m[i][j] = 0 或者 = 2就不必访问了
		return
	}
	m[i][j] = 2
	infect(m, line, col, i+1, j) // 感染右侧
	infect(m, line, col, i-1, j) // 感染左侧
	infect(m, line, col, i, j-1) // 感染上侧
	infect(m, line, col, i, j+1) // 感染下侧

}
