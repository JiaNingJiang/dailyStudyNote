package main

import "fmt"

func main() {
	str := "9,#,92,#,#"

	fmt.Println(isValidSerialization(str))
}

func isValidSerialization(preorder string) bool {
	n := len(preorder) // 字符串的总长度(包括分隔符',')

	stack := []int{1} // 模拟一个栈。初始栈中只有一个元素,值 == 1是因为提供给根节点使用的(根节点只有一个)

	for i := 0; i < n; i++ {
		if len(stack) == 0 { // 在遍历完前序字符串前,栈不能为空
			return false
		}

		if preorder[i] == ',' { // 当前字符为分隔符，跳过即可
			continue
		} else if preorder[i] == '#' { // 表示空节点，消耗栈顶元素一个槽位
			stack[len(stack)-1]--

			if stack[len(stack)-1] == 0 { // 栈顶元素槽位 == 0,槽位消耗殆尽,弹出该元素
				stack = stack[:len(stack)-1]
			}
		} else { // 表示正常节点.消耗栈顶元素一个槽位,同时栈中新加一个新的元素(有两个槽位)
			stack[len(stack)-1]--
			if stack[len(stack)-1] == 0 {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, 2) // 栈顶新加一个有两个槽位的元素
		}
	}

	return len(stack) == 0 // 遍历完前序字符串之后,栈需要为空
}
