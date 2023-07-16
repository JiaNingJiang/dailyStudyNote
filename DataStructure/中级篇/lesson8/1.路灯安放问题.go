package lesson8

func MinLight(area string) int {
	index := 0
	minLight := 0
	for {
		if index >= len(area) {
			return minLight
		}
		if area[index] == 'x' { // 当前位置是路障,则直接跳过
			index++
		} else { // 当前位置是人行道
			minLight++                // 该区域必然会有一个灯(具体位置可以是index或者是index+1，取决于下一个位置是否也是人行道)
			if index+1 == len(area) { // 当前位置就是最后一块人行道，则路灯就安放在当前index位置
				return minLight
			}
			if area[index+1] == 'x' { // 下一个位置是路障，则路灯就安放在当前index位置，然后跳到index+2位置
				index += 2
				continue
			} else { // 下一位置还是人行道，则路灯安放在index+1位置，然后跳到index+3位置
				index += 3
			}
		}
	}
}
