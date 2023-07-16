package lesson9

import "sort"

// 每一条活动路线起点任意，但是终点必须是最后一个活动
type ActRoute struct {
	TotalConsume int // 活动路线总计消耗的天数
	TotalSalary  int // 活动路线总计获得的报酬
}

type Activity struct {
	Consume  int         // 参加活动消耗的天数
	Salary   int         // 参加活动获得的报酬
	NextSet  []*Activity // 保存当前活动的下一个可以的活动
	LastSet  []*Activity // 用于追溯当前活动的上一个可以的活动
	OrderSet []*ActRoute // 活动路线的有序表(按照消耗天数从小到大有序，且天数与报酬需要成正比)
}

// 构建一个新活动
func NewActivity(consume, salary int, next, last []*Activity) *Activity {
	return &Activity{
		Consume:  consume,
		Salary:   salary,
		NextSet:  next,
		LastSet:  last,
		OrderSet: make([]*ActRoute, 0),
	}
}

func (act *Activity) AddNext(next *Activity) {
	act.NextSet = append(act.NextSet, next)
}

func (act *Activity) AddLast(next *Activity) {
	act.LastSet = append(act.LastSet, next)
}

// 整理指定节点为起点的活动路线
func (act *Activity) OrderActRoute() {

	if len(act.NextSet) == 0 {
		newRoute := &ActRoute{TotalConsume: act.Consume, TotalSalary: act.Salary}
		act.OrderSet = append(act.OrderSet, newRoute)
		return
	}

	for _, nextAct := range act.NextSet { // 遍历当前活动所有的后续活动
		for _, actRoute := range nextAct.OrderSet { // 遍历每一个后续活动的有序活动路线
			newRoute := &ActRoute{TotalConsume: act.Consume + actRoute.TotalConsume,
				TotalSalary: act.Salary + actRoute.TotalSalary}
			act.OrderSet = append(act.OrderSet, newRoute)
		}
	}

	// 将所有活动路线按照消耗天数进行排序
	sort.Slice(act.OrderSet, func(i, j int) bool {
		if act.OrderSet[i].TotalConsume < act.OrderSet[j].TotalConsume {
			return true
		} else {
			return false
		}
	})

	if len(act.OrderSet) == 1 {
		return
	}

	// 排除掉所有天数增多，但报酬不递增的活动路线
	targetRoutes := make([]*ActRoute, 0)
	index := 0
	for i := 0; i < len(act.OrderSet); i++ {
		if len(targetRoutes) == 0 {
			targetRoutes = append(targetRoutes, act.OrderSet[0])
			continue
		}
		if act.OrderSet[i].TotalSalary > targetRoutes[index].TotalSalary {
			targetRoutes = append(targetRoutes, act.OrderSet[i])
			index++
		}
	}
}

func SumPerfectRoutes(starts []*Activity) []*ActRoute {
	allActRoutes := make([]*ActRoute, 0)
	for _, start := range starts { // 遍历每一个活动
		for i := 0; i < len(start.OrderSet); i++ { // 遍历每一个活动的有效活动路线
			allActRoutes = append(allActRoutes, start.OrderSet...)
		}
	}

	sort.Slice(allActRoutes, func(i, j int) bool {
		if allActRoutes[i].TotalConsume < allActRoutes[j].TotalConsume {
			return true
		} else {
			return false
		}
	})

	targetRoutes := make([]*ActRoute, 0)
	index := 0
	for i := 0; i < len(allActRoutes); i++ {
		if len(targetRoutes) == 0 {
			targetRoutes = append(targetRoutes, allActRoutes[0])
			continue
		}
		if allActRoutes[i].TotalSalary > targetRoutes[index].TotalSalary {
			targetRoutes = append(targetRoutes, allActRoutes[i])
			index++
		}
	}

	return targetRoutes
}
