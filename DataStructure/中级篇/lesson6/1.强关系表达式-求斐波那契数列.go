package lesson6

import "DataStructure2/utils"

func selfMulMatrix(self interface{}, factor interface{}) interface{} {

	if self == nil { // self总是一个2x2的矩阵（开始是一个单位矩阵(对角线为1，其余位置为0)）
		selfMatrix := make([][]int, 2)
		selfMatrix[0] = make([]int, 2)
		selfMatrix[1] = make([]int, 2)
		selfMatrix[0][0] = 1
		selfMatrix[0][1] = 0
		selfMatrix[1][0] = 0
		selfMatrix[1][1] = 1
		factorMatrix := factor.([][]int)
		return utils.MatrixMultiply(selfMatrix, factorMatrix)
	} else {
		selfMatrix := self.([][]int)
		factorMatrix := factor.([][]int)
		return utils.MatrixMultiply(selfMatrix, factorMatrix)
	}
}

func Fibonacci(n int64) int {
	if n <= 2 {
		if n == 0 {
			return 0
		} else if n == 1 {
			return 1
		} else if n == 2 {
			return 1
		} else {
			panic("错误的输入")
		}
	}

	factorMatrix := [][]int{{1, 1}, {1, 0}} // 系数矩阵

	factorMatrixPow := EffectivePow(factorMatrix, n-2, selfMulMatrix) // 求系数矩阵的 n-2次幂

	initial := [][]int{{1, 1}}
	res := utils.MatrixMultiply(initial, factorMatrixPow.([][]int))

	return res[0][0]
}
