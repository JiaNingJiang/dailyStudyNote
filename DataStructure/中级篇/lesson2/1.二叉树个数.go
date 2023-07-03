package lesson2

// 给你n个节点，返回可以组成的二叉树的个数
func BinaryTreeCount(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	count := 0
	for left := 0; left <= n-1; left++ { // 除了根节点外，实际可控制的节点个数为n-1个
		right := n - 1 - left
		leftCount := BinaryTreeCount(left)
		rightCount := BinaryTreeCount(right)

		if left == 0 {
			count += rightCount
		} else if right == 0 {
			count += leftCount
		} else if left != 0 && right != 0 {
			count += leftCount * rightCount
		}
	}
	return count
}

// 用动态规划实现
func BinaryTreeCount2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	if n <= 2 {
		return dp[n]
	}

	for node := 3; node <= n; node++ { // 节点的个数
		for left := 0; left <= node-1; left++ {
			right := node - 1 - left
			if left == 0 {
				dp[node] += dp[right]
			} else if right == 0 {
				dp[node] += dp[left]
			} else if left != 0 && right != 0 {
				dp[node] += dp[left] * dp[right]
			}
		}
	}

	return dp[n]
}
