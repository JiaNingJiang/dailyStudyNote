package nonComparativeSort

func CountingSort(arr []int) {
	var maxValue int = 100              // 测试用数组中每个元素的最大值
	sortArr := make([]int, 0, len(arr)) // 存储最终排序完成的数组(升序)
	countArr := make([]int, maxValue)   // 计数用哈希表
	for _, v := range arr {             // 遍历原始数组，每遇到一个数字，将其在hash表中对应的出现次数+1
		countArr[v]++
	}
	for num, count := range countArr { // num表示一个元素，count表示该元素在arr数组中出现的次数
		if count <= 0 {
			continue
		} else {
			for j := 0; j < count; j++ { // 将这count个num追加到sortArr中
				sortArr = append(sortArr, num)
			}
		}
	}

	for i, v := range sortArr {
		arr[i] = v
	}

}
