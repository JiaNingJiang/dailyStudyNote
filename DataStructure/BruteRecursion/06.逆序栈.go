package BruteRecursion

import "DataStructure/linkList"

func ReverseStack(stack *linkList.Stack) {
	if stack.Len == 0 { // 栈中所有元素被弹出时，开始返回
		return
	}

	bottom := getBottom(stack) // 弹出栈的最底层元素
	ReverseStack(stack)
	stack.Push(bottom) // 因为递归进行返回时是相反的顺序，因此将栈中元素逆序填入到栈中
}

// 弹出栈底元素
func getBottom(stack *linkList.Stack) interface{} {
	result := stack.Pop() // 将当前栈顶元素弹出
	if stack.Len == 0 {   // 当遇到栈底元素时，不会再次加入到栈中，而是直接返回
		return result
	}
	last := getBottom(stack) // 不断请求获取下一个元素
	stack.Push(result)       // 重新将当前栈顶元素加入
	return last              // 从最后一次递归获得栈底元素
}
