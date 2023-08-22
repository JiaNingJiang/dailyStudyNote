package greedy

import "math"

func NQueen(n int) int {
	if n < 1 { // n == 0，解法为0
		return 0
	}
	chessBoard := make([]int, n, n) // 棋盘。 chessBoard[i]的取值范围是 k = 0~n-1 ，表示第i行当前在第k列
	return nQueen(0, n, chessBoard)
}

// line:当前的行号，n棋盘总行书，chessBoard数组模拟的棋盘
func nQueen(line, n int, chessBoard []int) int {
	if line == n { // 0~n-1行完成了N皇后问题，返回解法数量+1
		return 1
	}
	result := 0
	// 当前第i行行的所有列都要进行尝试
	for col := 0; col < n; col++ {
		// 1.先判断当前第line行的col列是否可以放皇后
		if isValid(chessBoard, line, col) { // 不会和前i-1行冲突，则放，并继续递归找下一行的放置
			chessBoard[line] = col                  // 放置皇后
			result += nQueen(line+1, n, chessBoard) //  完成验证，继续摆放第i+1行
		}
	}
	return result
}

func isValid(chessBoard []int, line, col int) bool {
	for i := 0; i < line; i++ { // 检查line,col放置皇后是否与前面line-1行有冲突
		if col == chessBoard[i] { // 共列
			return false
		}
		if line-i == int(math.Abs(float64(col-chessBoard[i]))) { // 列数之差等于行数之差，共斜线
			return false
		}
	}
	return true // 检查完成，没有发现任何冲突
}
