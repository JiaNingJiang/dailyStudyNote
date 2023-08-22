package monotonicStack

import (
	"DataStructure/08.linkList"
	"reflect"
)

type Element struct {
	value interface{} // 元素本身的值
	index int         // 元素在原数组中的下标
}

type MonotonicStack struct {
	baseStack      *linkList.Stack // 底层栈
	unorderedArray []interface{}   // 待排序的底层数组
	form           bool            // form == true则为单调递增栈(从栈顶往栈底看，是单调递增的关系);否则为单调递减栈
}

func NewMonotonicStack(src interface{}, form bool) *MonotonicStack {
	mStack := new(MonotonicStack)
	mStack.baseStack = linkList.NewStack()

	arr := src.([]int)
	for _, v := range arr {
		mStack.unorderedArray = append(mStack.unorderedArray, v)
	}
	mStack.form = form
	return mStack
}

// 适用于没有重复值的单调栈
// 对整个底层数组进行遍历,数组的每个元素都会获得其左右两侧最近的比他大(或者小)的两个元素
func (mStack *MonotonicStack) Order() map[interface{}][2]interface{} {

	result := make(map[interface{}][2]interface{}, 0) // 返回结果，key为一个元素，value分别包含该元素左右两侧距离最近的比他大(小)的元素

	mStack.baseStack = linkList.NewStack()

	index := 0
	if mStack.form { // 单调递增栈
		for {
			if mStack.baseStack.Len == 0 { // 栈被清空,直接将当前元素的记录存入栈中
				mStack.baseStack.Push(index)
				index++
				continue
			}

			if index == len(mStack.unorderedArray) { // 没有剩余元素可以入栈,则结束
				for { // 将栈中剩余元素按次序弹出
					if mStack.baseStack.Len == 0 {
						return result
					}
					eleIndex := mStack.baseStack.Pop()           // 元素的下标
					ele := mStack.unorderedArray[eleIndex.(int)] // 元素本身
					leftBigIndex := mStack.baseStack.Top()
					var leftBig interface{}
					var leftEle Element
					if leftBigIndex == nil {
						leftBig = nil
						leftEle = Element{}
					} else {
						leftBig = mStack.unorderedArray[leftBigIndex.(int)]
						leftEle = Element{value: leftBig, index: leftBigIndex.(int)}
					}
					if reflect.DeepEqual(leftEle, Element{}) {
						result[ele] = [2]interface{}{nil, nil} // 左侧是其在栈中的下一个元素，右侧没有
					} else {
						result[ele] = [2]interface{}{leftEle, nil} // 左侧是其在栈中的下一个元素，右侧没有
					}

				}

			}
			currentEle := mStack.unorderedArray[index]                    // 当前元素
			topEle := mStack.unorderedArray[mStack.baseStack.Top().(int)] // 栈顶所记录的元素
			if lessElement(currentEle, topEle) {                          // 当前需要进栈的元素小于栈顶元素，则可以直接将记录入栈成为新的栈顶
				mStack.baseStack.Push(index)
				index++
			} else { // 当前需要进栈的元素大于栈顶元素，则需要不断将旧的元素弹出，直到需要进栈的元素小于栈顶元素
				eleIndex := mStack.baseStack.Pop()           // 元素的下标
				ele := mStack.unorderedArray[eleIndex.(int)] // 元素本身
				leftBigIndex := mStack.baseStack.Top()
				var leftBig interface{}
				var leftEle Element
				if leftBigIndex == nil {
					leftBig = nil
					leftEle = Element{}
				} else {
					leftBig = mStack.unorderedArray[leftBigIndex.(int)]
					leftEle = Element{value: leftBig, index: leftBigIndex.(int)}
				}
				rightEle := Element{value: currentEle, index: index}
				if reflect.DeepEqual(leftEle, Element{}) {
					result[ele] = [2]interface{}{nil, rightEle}
				} else {
					result[ele] = [2]interface{}{leftEle, rightEle} // 左侧是其在栈中的下一个元素，右侧是当前需要进栈的元素
				}

				// index不变,这样做的目的是不断将旧的元素弹出，直到需要进栈的元素小于栈顶元素或者栈清空
			}
		}
	} else { // 单调递减栈
		for {
			if mStack.baseStack.Len == 0 { // 栈被清空,直接将当前元素的记录存入栈中
				mStack.baseStack.Push(index)
				index++
				continue
			}

			if index == len(mStack.unorderedArray) { // 没有剩余元素可以入栈,则结束
				for { // 将栈中剩余元素按次序弹出
					if mStack.baseStack.Len == 0 {
						return result
					}
					eleIndex := mStack.baseStack.Pop()           // 元素的下标
					ele := mStack.unorderedArray[eleIndex.(int)] // 元素本身
					leftSmallIndex := mStack.baseStack.Top()
					var leftSmall interface{}
					var leftEle Element
					if leftSmallIndex == nil {
						leftSmall = nil
						leftEle = Element{}
					} else {
						leftSmall = mStack.unorderedArray[leftSmallIndex.(int)]
						leftEle = Element{value: leftSmall, index: leftSmallIndex.(int)}
					}
					if reflect.DeepEqual(leftEle, Element{}) {
						result[ele] = [2]interface{}{nil, nil} // 左侧是其在栈中的下一个元素，右侧没有
					} else {
						result[ele] = [2]interface{}{leftEle, nil} // 左侧是其在栈中的下一个元素，右侧没有
					}
				}
			}

			currentEle := mStack.unorderedArray[index]                    // 当前元素
			topEle := mStack.unorderedArray[mStack.baseStack.Top().(int)] // 栈顶所记录的元素
			if !lessElement(currentEle, topEle) {                         // 当前需要进栈的元素大于栈顶元素，则可以直接将记录入栈成为新的栈顶
				mStack.baseStack.Push(index)
				index++
			} else { // 当前需要进栈的元素小于栈顶元素，则需要不断将旧的元素弹出，直到需要进栈的元素大于栈顶元素
				eleIndex := mStack.baseStack.Pop()           // 元素的下标
				ele := mStack.unorderedArray[eleIndex.(int)] // 元素本身
				leftSmallIndex := mStack.baseStack.Top()
				var leftSmall interface{}
				var leftEle Element
				if leftSmallIndex == nil {
					leftSmall = nil
					leftEle = Element{}
				} else {
					leftSmall = mStack.unorderedArray[leftSmallIndex.(int)]
					leftEle = Element{value: leftSmall, index: leftSmallIndex.(int)}
				}
				rightEle := Element{value: currentEle, index: index}
				if reflect.DeepEqual(leftEle, Element{}) {
					result[ele] = [2]interface{}{nil, rightEle}
				} else {
					result[ele] = [2]interface{}{leftEle, rightEle} // 左侧是其在栈中的下一个元素，右侧是当前需要进栈的元素
				}
				// index不变,这样做的目的是不断将旧的元素弹出，直到需要进栈的元素大于栈顶元素或者栈清空
			}
		}
	}
}

// 适用于有重复值的单调栈
// 对整个底层数组进行遍历,数组的每个元素都会获得其左右两侧最近的比他大(或者小)的两个元素
func (mStack *MonotonicStack) Order2() map[interface{}][2]interface{} {
	result := make(map[interface{}][2]interface{}, 0) // 返回结果，key为一个元素，value分别包含该元素左右两侧距离最近的比他大(小)的元素
	mStack.baseStack = linkList.NewStack()

	index := 0
	if mStack.form { // 单调递增栈
		for {
			if mStack.baseStack.Len == 0 { // 栈被清空,将当前元素作为链表头入栈
				link := linkList.NewLinkList()
				link.TailAdd(index)
				mStack.baseStack.Push(link)
				index++
				continue
			}

			if index == len(mStack.unorderedArray) { // 没有剩余元素可以入栈,则结束
				for { // 将栈中剩余元素按次序弹出
					if mStack.baseStack.Len == 0 {
						return result
					}
					link := mStack.baseStack.Pop().(*linkList.LinkList) // 弹出存储下标的整条链表

					// 一条链表内指代的元素的大小都是一样的,因此其左右侧值是一样的ll.Len--
					leftSmallIndex := mStack.baseStack.Top()
					var leftSmall interface{}
					var leftEle Element
					if leftSmallIndex == nil {
						leftSmall = nil
						leftEle = Element{}
					} else {
						leftSmall = mStack.unorderedArray[leftSmallIndex.(*linkList.LinkList).Head.Data.(int)]
						leftEle = Element{value: leftSmall, index: leftSmallIndex.(*linkList.LinkList).Head.Data.(int)}
					}

					for { // 按照从链表尾部到头部的顺序相继弹出
						if link.Len == 0 { // 当前链表中的所有元素都已经完成弹出
							break
						}
						eleIndex := link.TailDel()                   // 元素下标
						ele := mStack.unorderedArray[eleIndex.(int)] // 元素本身

						key := Element{value: ele, index: eleIndex.(int)}
						if reflect.DeepEqual(leftEle, Element{}) {
							result[key] = [2]interface{}{nil, nil}
						} else {
							result[key] = [2]interface{}{leftEle, nil}
						}
					}
				}
			}

			currentEle := mStack.unorderedArray[index] // 当前要入栈的元素大小
			topIndex := mStack.baseStack.Top().(*linkList.LinkList).Head.Data
			topEle := mStack.unorderedArray[topIndex.(int)] // 栈顶所记录的元素大小

			cmp := compareElement(currentEle, topEle)

			if cmp == -1 { // 当前需要进栈的元素小于栈顶元素，新建新的链表作为头部，入栈
				link := linkList.NewLinkList()
				link.TailAdd(index)
				mStack.baseStack.Push(link)
				index++
			} else if cmp == 0 { // 等于栈顶元素,将其插入到栈顶元素链表中
				mStack.baseStack.Top().(*linkList.LinkList).TailAdd(index)
				index++
			} else { // 大于栈顶元素,需要不断将旧的元素弹出，直到需要进栈的元素小于等于栈顶元素
				link := mStack.baseStack.Pop().(*linkList.LinkList) // 弹出存储下标的整条链表
				// 一条链表内指代的元素的大小都是一样的,因此其左右侧值是一样的
				leftSmallIndex := mStack.baseStack.Top()
				var leftSmall interface{}
				var leftEle Element
				if leftSmallIndex == nil {
					leftSmall = nil
					leftEle = Element{}
				} else {
					leftSmall = mStack.unorderedArray[leftSmallIndex.(*linkList.LinkList).Head.Data.(int)]
					leftEle = Element{value: leftSmall, index: leftSmallIndex.(*linkList.LinkList).Head.Data.(int)}
				}
				rightEle := Element{value: currentEle, index: index}

				for { // 按照从链表尾部到头部的顺序相继弹出
					if link.Len == 0 { // 当前链表中的所有元素都已经完成弹出
						break
					}
					eleIndex := link.TailDel()                   // 元素下标
					ele := mStack.unorderedArray[eleIndex.(int)] // 元素本身
					key := Element{value: ele, index: eleIndex.(int)}

					if reflect.DeepEqual(leftEle, Element{}) {
						result[key] = [2]interface{}{nil, rightEle}
					} else {
						result[key] = [2]interface{}{leftEle, rightEle}
					}

				}
				// index不变,这样做的目的是不断将旧的元素弹出，直到需要进栈的元素大于栈顶元素或者栈清空
			}

		}

	} else {
		panic("to do implement")
	}

}

func lessElement(current, top interface{}) bool {
	if current.(int) < top.(int) {
		return true
	} else {
		return false
	}
}

// 比较两个元素的大小： > 返回1，==返回0，< 返回-1
func compareElement(current, top interface{}) int {
	if current.(int) > top.(int) {
		return 1
	} else if current.(int) == top.(int) {
		return 0
	} else {
		return -1
	}
}
