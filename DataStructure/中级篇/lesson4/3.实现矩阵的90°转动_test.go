package lesson4

import "testing"

func TestNinetyRotation(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 3},
		{4, 5, 6, 7},
		{8, 9, 10, 11},
		{12, 13, 14, 15},
	}
	NinetyRotation(&matrix)

	PrintMatrix(matrix)
}
