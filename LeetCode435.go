package main

// 求所有区间中不重叠区间的最多个数
func eraseOverlapIntervals(intervals [][]int) int {
	quickSort(intervals, 0, len(intervals)-1)
	endpos := intervals[0][1]
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= endpos {
			count++
			endpos = intervals[i][1]
			continue
		} else {
			//// 右0大于左1
			//if intervals[i][1] >= endpos {
			//	endpos = intervals[i][1]
			//}
		}
	}
	return len(intervals) - count
}

func quickSort(intervals [][]int, low int, high int) {
	if low >= high {
		return
	}
	partition := partition(intervals, low, high)
	quickSort(intervals, low, partition-1)
	quickSort(intervals, partition+1, high)
}

func partition(intervals [][]int, low int, high int) int {
	i := low
	j := high
	for i < j {
		for i < j && intervals[j][1] >= intervals[low][1] {
			j--
		}
		for i < j && intervals[i][1] <= intervals[low][1] {
			i++
		}
		intervals[i], intervals[j] = intervals[j], intervals[i]
	}
	intervals[i], intervals[low] = intervals[low], intervals[i]
	return i
}
