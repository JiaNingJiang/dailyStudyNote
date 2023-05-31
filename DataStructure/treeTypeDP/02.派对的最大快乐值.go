package treeTypeDP

type staff struct {
	happyDegree int      // 本员工自己的快乐值
	nexts       []*staff // 本员工的下属
}

type happyInfo struct {
	laiHappyDegree   int // 当前员工来时的最大快乐值
	bulaiHappyDegree int // 当前员工不来是的最大快乐值
}

func GetMaxHappyDegree(boss *staff) int {
	info := getMaxHappyDegree(boss)
	return getMax(info.laiHappyDegree, info.bulaiHappyDegree)
}

func getMaxHappyDegree(root *staff) happyInfo {
	if root.nexts == nil { // 当前是基层员工
		return happyInfo{laiHappyDegree: root.happyDegree, bulaiHappyDegree: 0}
	}
	lai := root.happyDegree
	bulai := 0

	for _, next := range root.nexts {
		nextInfo := getMaxHappyDegree(next)                                 // 来自下属员工的返回信息
		lai += nextInfo.bulaiHappyDegree                                    // 当前员工已经来了，那么下属员工必定不能来
		bulai += getMax(nextInfo.laiHappyDegree, nextInfo.bulaiHappyDegree) // 当前员工不来，那么下属员工可以来也可以不来。这里选择快乐值最大的分支
	}
	return happyInfo{laiHappyDegree: lai, bulaiHappyDegree: bulai}
}
