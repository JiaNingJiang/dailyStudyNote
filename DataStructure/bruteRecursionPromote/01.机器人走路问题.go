package bruteRecursionPromote

import "fmt"

// N与E是固定值, N表示需要走的总步数   E表示目标地点
// left与right也是固定值，分别表示左边界与右边界
// 返回值为走法个数
func RobotWalk(N, E int, left, right int) int {
	if E < left || E > right {
		fmt.Printf("目标位置%d不在范围(%d ~ %d)内\n", E, left, right)
		return 0
	}
	return robotWalk(N, E, N, left, left, right)
}

// res表示剩余需要走的步数
// cur表示当前位置
func robotWalk(N, E, res, cur int, left, right int) int {
	if res == 0 { // 走完全程
		if cur == E { // 当前位置在E上,表示此走法可行
			return 1
		} else {
			return 0
		}
	}
	if cur == left { // 当前在左边界位置的话，下一步只能往右走
		return robotWalk(N, E, res-1, left+1, left, right)
	}
	if cur == right { // 当前在右边界位置的话，下一步只能往左走
		return robotWalk(N, E, res-1, right-1, left, right)
	}
	return robotWalk(N, E, res-1, cur-1, left, right) + robotWalk(N, E, res-1, cur+1, left, right) // 正常情况下，分左右两种可能
}

// 暴力递归+记忆化缓存
func RobotWalkMemoryCache(N, E int, left, right int) int {
	if E < left || E > right {
		fmt.Printf("目标位置%d不在范围(%d ~ %d)内\n", E, left, right)
		return 0
	}

	dpMemory := make([][]int, N+1) // dpMemory即是记忆化缓存
	for i := 0; i <= N; i++ {      // 横坐标表示当前剩余步数，总是在0~N之间
		dpMemory[i] = make([]int, right+1)
		for j := left; j <= right; j++ { // 纵坐标表示当前位置，总是在left~right之间
			dpMemory[i][j] = -1
		}
	}
	return robotWalkMemoryCache(N, E, dpMemory, N, left, left, right)
}

func robotWalkMemoryCache(N, E int, dpMemory [][]int, res int, cur int, left, right int) int {
	//if res < 0 || res > N {
	//	fmt.Println("横坐标越界访问....,res = ", res)
	//	return -1
	//}
	//if cur < left || cur > right {
	//	fmt.Printf("纵坐标(%d ~ %d)越界访问....,cur = %d \n", left, right, cur)
	//	return -1
	//}
	if dpMemory[res][cur] != -1 { // 缓存命中( 为1 或者 为0)
		return dpMemory[res][cur]
	}
	if res == 0 { // 走完全程，1表示此走法可行，0表示此走法不可行
		if cur == E {
			dpMemory[res][cur] = 1
		} else {
			dpMemory[res][cur] = 0
		}
		return dpMemory[res][cur]
	}
	if cur == left { // 当前在左边界,下一步只能向右走
		dpMemory[res][cur] = robotWalkMemoryCache(N, E, dpMemory, res-1, left+1, left, right)
	} else if cur == right { // 当前在右边界,下一步只能向左走
		dpMemory[res][cur] = robotWalkMemoryCache(N, E, dpMemory, res-1, right-1, left, right)
	} else { // 普通情况，分为向左、向右两种情况
		dpMemory[res][cur] = robotWalkMemoryCache(N, E, dpMemory, res-1, cur+1, left, right) +
			robotWalkMemoryCache(N, E, dpMemory, res-1, cur-1, left, right)
	}
	return dpMemory[res][cur]
}

func RobotWalk3(N, E int, left, right int) int {
	if E < left || E > right {
		fmt.Printf("目标位置%d不在范围(%d ~ %d)内\n", E, left, right)
		return 0
	}
	table := make([][]int, N+1) // 严格表结构：横坐标表示当前剩余步数，范围是0~N；纵坐标表示当前位置，范围是0~right之间
	for i := 0; i <= N; i++ {   //
		table[i] = make([]int, right+1)
	}

	// 根据初始条件得到第0行的数据(第零行只有(0,E)=1，其余都为0)  (left列左侧的全部列不应参与)
	for col := left; col <= right; col++ {
		if col == E {
			table[0][col] = 1
		} else {
			table[0][col] = 0
		}
	}
	for line := 1; line <= N; line++ { // 从第1行到第N行进行迭代  (left列左侧的全部列不应参与)
		for col := left; col <= right; col++ {
			if col == left { // 特殊：当前在left位置，那么具体的路径数只取决于left+1位置
				table[line][left] = table[line-1][left+1]
				continue
			}
			if col == right { // 特殊：当前在right位置，具体的路径数只取决于right-1位置
				table[line][right] = table[line-1][right-1]
				continue
			}
			// 正常情况，具体的路径数同时取决于col-1和col+1位置
			table[line][col] = table[line-1][col-1] + table[line-1][col+1]
		}
	}
	return table[N][left]
}
