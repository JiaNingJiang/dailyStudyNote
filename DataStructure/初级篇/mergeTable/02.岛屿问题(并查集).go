package mergeTable

import "sync"

// 描述矩阵中一个点的坐标
type point struct {
	line int
	col  int
}

type Chart struct {
	baseMatrix      [][]int // 底层矩阵
	baseMatrixMutex sync.Mutex
	line            int // 矩阵总行数
	col             int // 矩阵总列数

	island      int // 岛屿数目
	islandMutex sync.Mutex

	unionFindSet      *UnionFindSet // 并查集
	unionFindSetMutex sync.Mutex
}

func NewChart(m [][]int) *Chart {
	if m == nil || m[0] == nil {
		return nil
	}

	chart := new(Chart)
	chart.baseMatrix = m
	chart.line = len(m)   // 矩阵行数
	chart.col = len(m[0]) // 矩阵列数

	coordinatePoint := make([]interface{}, 0) // 存储矩阵中所有值为1的节点的坐标

	for i := 0; i < chart.line; i++ {
		for j := 0; j < chart.col; j++ {
			if m[i][j] == 1 {
				p := point{line: i, col: j}
				coordinatePoint = append(coordinatePoint, p)
			}
		}
	}

	chart.unionFindSet = NewUnionFindSet(coordinatePoint) // 初始化：每一个值为1的拥有一个集合

	return chart
}

func (chart *Chart) CountIsland() {
	// 分成左右两部分进行并行查询
	split := chart.col / 2
	var wg sync.WaitGroup
	wg.Add(2)
	go chart.partitionQuery(0, split-1, 0, chart.line-1, &wg)         // 负责查询左半区域
	go chart.partitionQuery(split, chart.col-1, 0, chart.line-1, &wg) // 负责查询右半区域

	wg.Wait()

	// 在分界线上进行合并
	for line := 0; line < chart.line; line++ {
		// 分界线左右都为2，可以尝试进行合并
		if chart.baseMatrix[line][split-1] == 2 && chart.baseMatrix[line][split] == 2 {
			leftPoint := point{line: line, col: split - 1}
			rightPoint := point{line: line, col: split}

			// 如果分界线左右不在同一个集合中(特征节点不一样)，需要进行合并
			if chart.unionFindSet.IsSameSet(leftPoint, rightPoint) {
				continue
			} else {
				chart.unionFindSet.Union(leftPoint, rightPoint)
				chart.island--
			}
		}

	}

}

// 负责在指定区域内查询岛屿的数量
func (chart *Chart) partitionQuery(left, right, upper, down int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := upper; i <= down; i++ {
		for j := left; j <= right; j++ {
			if chart.baseMatrix[i][j] == 1 {
				chart.islandMutex.Lock()
				chart.island++
				chart.islandMutex.Unlock()

				p := point{line: i, col: j} // 当前节点作为本组的特征节点，持续合并同组的其他节点
				chart.infect(i, j, left, right, upper, down, p)
			}
		}
	}
}

func (chart *Chart) infect(line, col int, left, right, upper, down int, p point) {
	if line < upper || line > down || col < left || col > right || chart.baseMatrix[line][col] != 1 { // 1.访问矩阵不能越界  2.如果m[i][j] = 0 或者 = 2就不必访问了
		return
	}
	chart.baseMatrixMutex.Lock()
	chart.baseMatrix[line][col] = 2
	chart.baseMatrixMutex.Unlock()

	current := point{line: line, col: col}
	chart.unionFindSetMutex.Lock()
	chart.unionFindSet.Union(p, current) // 将当前节点与本组的特征节点进行集合合并
	chart.unionFindSetMutex.Unlock()

	chart.infect(line, col-1, left, right, upper, down, p) // 感染左侧
	chart.infect(line, col+1, left, right, upper, down, p) // 感染右侧
	chart.infect(line-1, col, left, right, upper, down, p) // 感染上侧
	chart.infect(line+1, col, left, right, upper, down, p) // 感染下侧
}
