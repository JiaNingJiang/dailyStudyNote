package linkList

import "DataStructure/05.quickSort"

func SeparateLinkList(ll *LinkList, pivot int) {
	if ll.Len < 2 {
		return
	}
	arr := make([]int, 0, ll.Len) // 存储链表数据
	current := ll.Head
	for { // 第一个循环，按照链表原来的顺序，将存储的数据移动到临时数组中
		if current == nil {
			break
		}
		data, _ := current.Data.(int)
		arr = append(arr, data)
		current = current.Next
	}
	quickSort.DutchFlag(arr, pivot) // 采用荷兰国旗分区方法: 左区域小于pivot，中间等于pivot，右边大于pivot
	current = ll.Head
	index := 0
	for { // 第二个循环，利用分区排序后的数组重新更新原始链表
		if current == nil {
			break
		}
		current.Data = arr[index]
		index++
		current = current.Next
	}
}
