package BruteRecursion

func PrintAllCombination(str string) (int, map[string]struct{}) {
	charSet := []byte(str)
	result := make(map[string]struct{}, 0)
	printAllCombination(charSet, 0, result)
	return len(result), result
}

// 存在一些重复的递归分支(因为字符串中可能会出现重复的字符)
func printAllCombination(str []byte, index int, res map[string]struct{}) {
	if index == len(str) {
		res[string(str)] = struct{}{}
	}

	for i := index; i < len(str); i++ {
		swap(str, index, i) // 将字符串的当前位替换为 index后的所有位
		printAllCombination(str, index+1, res)
		swap(str, index, i) // 将递归分支分离出去后再交换回去
	}
}

func PrintAllCombinationImproved(str string) (int, map[string]struct{}) {
	charSet := []byte(str)
	result := make(map[string]struct{}, 0)
	printAllCombinationImproved(charSet, 0, result)
	return len(result), result
}

func printAllCombinationImproved(str []byte, index int, res map[string]struct{}) {
	if index == len(str) {
		res[string(str)] = struct{}{}
	}
	visitMap := make(map[byte]bool)
	var a byte = 97
	var A byte = 65
	for i := 0; i < 26; i++ {
		visitMap[a+byte(i)] = false
		visitMap[A+byte(i)] = false
	}

	for i := index; i < len(str); i++ {
		if !visitMap[str[i]] {
			visitMap[str[i]] = true
			swap(str, index, i) // 将字符串的当前位替换为 index后的所有位
			printAllCombination(str, index+1, res)
			swap(str, index, i) // 将递归分支分离出去后再交换回去
		}

	}
}

func swap(arr []byte, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
