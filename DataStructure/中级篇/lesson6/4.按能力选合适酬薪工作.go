package lesson6

import (
	"sort"
)

type Job struct {
	Ability int // 工作需要的能力
	Salary  int // 工作提供的薪资
}

// 工作数组jobArr，个人能力数组capArr
func AcAbilityGetWork(jobArr []Job, capArr []int) []int {
	// 1.将工作数组按照工作难度从小到大排序
	sort.Slice(jobArr, func(i, j int) bool {
		if jobArr[i].Ability < jobArr[j].Ability {
			return true
		} else {
			return false
		}
	})

	filterArr := make([]Job, 0) // 只记录各工作难度下，薪资最高的工作
	maxSalary := jobArr[0].Salary
	curCap := jobArr[0].Ability
	// 2.对于工作难度相等的工作，jobArr中只保留薪资最高的那一个
	for i := 1; i < len(jobArr); i++ {
		if jobArr[i].Ability == curCap { // 工作难度相等
			if jobArr[i].Salary > maxSalary {
				maxSalary = jobArr[i].Salary // 在相同的工作难度下，记录薪资最高的一个
			} else {
				if i == len(jobArr)-1 {
					filterArr = append(filterArr, Job{Salary: maxSalary, Ability: curCap})
				}
			}
		} else { // 只能是工作难度增大了，而不可能是变小了
			filterArr = append(filterArr, Job{Salary: maxSalary, Ability: curCap}) // 记录上一个工作难度下的，薪资最高的工作
			maxSalary = jobArr[i].Salary
			curCap = jobArr[i].Ability
		}
	}

	increaseJob := make([]Job, 0) // 工作如果难度上升，那么薪资也需要提高

	increaseJob = append(increaseJob, filterArr[0])
	lastJob := 0
	for i := 1; i < len(filterArr); i++ {
		if filterArr[i].Salary >= increaseJob[lastJob].Salary {
			increaseJob = append(increaseJob, filterArr[i])
			lastJob++
		}
	}

	res := make([]int, 0)
	for i := 0; i < len(capArr); i++ {
		personCap := capArr[i]
		salary := 0
		for _, job := range increaseJob {
			if personCap >= job.Ability {
				salary = job.Salary
			}
		}
		res = append(res, salary)
	}

	return res
}
