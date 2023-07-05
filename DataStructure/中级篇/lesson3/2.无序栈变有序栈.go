package lesson3

import (
	"DataStructure2/utils"
)

// increase == true表示变成递增站；increase == false表示变成递减栈
func Orderization(stack *utils.Stack, increase bool) {
	aStack := utils.NewStack() // 辅助栈

	// 将原始栈中的元素全部移动到辅助栈，同时保证辅助栈与原始栈的目标单调性相反
	for {
		if stack.Len == 0 {
			break
		}
		ele := stack.Pop().(int) // 从原始栈中弹出一个元素

		if aStack.Len == 0 { // 仅当辅助栈为空时可以直接加入到辅助栈
			aStack.Push(ele)
			continue
		}
		top := aStack.Top().(int)
		if compare(ele, top, !increase) { // 需要确保辅助栈的栈顶元素与原始栈中弹出的元素符合相反的单调关系
			aStack.Push(ele)
		} else { // 如果不满足单调关系，则需要从辅助栈中持续弹出元素，直到符合单调关系或者辅助栈为空
			for {
				if aStack.Len == 0 {
					aStack.Push(ele)
					break
				}
				top := aStack.Top().(int)
				if compare(ele, top, !increase) {
					aStack.Push(ele)
					break
				}
				data := aStack.Pop()
				stack.Push(data)
			}
		}
	}

	for {
		if aStack.Len == 0 {
			return
		}
		data := aStack.Pop()
		stack.Push(data)
	}

}

func compare(original, assistTop int, increase bool) bool {
	if increase {
		if original > assistTop {
			return true
		} else {
			return false
		}
	} else {
		if original < assistTop {
			return true
		} else {
			return false
		}
	}
}
