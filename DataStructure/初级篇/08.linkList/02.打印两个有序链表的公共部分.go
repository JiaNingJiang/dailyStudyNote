package linkList

import "fmt"

const (
	TypeInt = iota
	TypeFloat
	TypeString
)

func PublicSection(l1, l2 *LinkList) {
	l1current := l1.Head
	l2current := l2.Head
	commonSection := make([]interface{}, 0)
	if l1current == nil || l2current == nil {
		fmt.Println("没有公共部分")
		return
	}
	for {
		if l1current == nil || l2current == nil {
			break
		}
		compare := NodeDataCompare(l1current.Data, l2current.Data, TypeInt)
		if compare == 0 {
			commonSection = append(commonSection, l1current.Data)
			l1current = l1current.Next
			l2current = l2current.Next
		} else if compare == -1 {
			l1current = l1current.Next
		} else if compare == 1 {
			l2current = l2current.Next
		} else {
			panic("链表节点数据段类型不匹配")
		}
	}

	fmt.Printf("公共部分为：%v", commonSection)
}

func NodeDataCompare(d1, d2 interface{}, dataType int) int {
	switch dataType {
	case TypeInt:
		data1, _ := d1.(int)
		data2, _ := d2.(int)
		if data1 == data2 {
			return 0
		} else if data1 < data2 {
			return -1
		} else {
			return 1
		}
	case TypeFloat:
		data1, _ := d1.(float32)
		data2, _ := d2.(float32)
		if data1 == data2 {
			return 0
		} else if data1 < data2 {
			return -1
		} else {
			return 1
		}
	case TypeString:
		data1, _ := d1.(string)
		data2, _ := d2.(string)
		if data1 == data2 {
			return 0
		} else if data1 < data2 {
			return -1
		} else {
			return 1
		}
	default:
		return 2
	}
}
