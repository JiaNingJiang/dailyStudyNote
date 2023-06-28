package greedy

import "sort"

type Meeting struct {
	Start int // 会议的开始时间
	End   int // 会议的结束时间
}

// meetings是所有有需求的会议，start是会议室最早可以开始使用的时间
// 返回值是安排上的会议集合
func BestArrage(meetings []Meeting, start int) []Meeting {
	sort.Slice(meetings, func(i, j int) bool { // 按照每场会议的End进行升序排序
		if meetings[i].End < meetings[j].End {
			return true
		} else {
			return false
		}
	})
	result := make([]Meeting, 0)
	for _, meeting := range meetings { // 遍历所有会议，如果当前会议在之前会议结束之后开始，则将其加入到结果集合中
		if meeting.Start >= start {
			result = append(result, meeting)
			start = meeting.End // 更新会议室的最新结束时间
		}
	}
	return result
}

func BestArrageForce(meetings []Meeting, start int) []Meeting {
	sort.Slice(meetings, func(i, j int) bool { // 按照每场会议的End进行升序排序（重要：暴力递归方法也需要先按照会议的end时间进行升序排序）
		if meetings[i].End < meetings[j].End {
			return true
		} else {
			return false
		}
	})
	selects := make([]Meeting, 0)
	return bestArrageForce(meetings, start, 0, selects)
}

// 暴力实现（一场会议，要么可以使用会议室，要么不能使用）
func bestArrageForce(meetings []Meeting, lastTime int, index int, selects []Meeting) []Meeting {
	if index >= len(meetings) { // 越界访问，不再向下递归，直接返回已选出的会议
		return selects
	}
	useSelect := make([]Meeting, 0)
	useSelect = append(useSelect, selects...)
	notUseSelect := make([]Meeting, 0)
	notUseSelect = append(notUseSelect, selects...)
	meeting := meetings[index] // 当前会议
	// 1.意图举行当前会议(可能会举行当前会议，取决于当前条件是否允许，即当前会议开始时间是否在最晚一次会议室结束时间之后)
	useResult := make([]Meeting, 0)
	if meeting.Start >= lastTime { // 条件允许，可以举行当前会议
		useSelect = append(useSelect, meeting)
		useResult = append(useResult, bestArrageForce(meetings, meeting.End, index+1, useSelect)...)
	} else { // 条件不允许，不可以举行当前会议
		useResult = append(useResult, bestArrageForce(meetings, lastTime, index+1, useSelect)...)
	}

	// 2.意图不举行当前会议(必定不会举行当前会议)
	notUseResult := make([]Meeting, 0)
	notUseResult = append(notUseResult, bestArrageForce(meetings, lastTime, index+1, notUseSelect)...)

	if len(useResult) > len(notUseResult) {
		return useResult
	} else {
		return notUseResult
	}
}
