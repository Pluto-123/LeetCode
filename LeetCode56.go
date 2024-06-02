package main

import "fmt"

/*
*
示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	// 先排序
	//sort.Slice(intervals, func(i, j int) bool {
	//	return intervals[i][0] < intervals[j][0]
	//})
	for i := 0; i < len(intervals); {
		left := intervals[i][0]
		right := intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] > intervals[i][1] {
				// 连不上
				i = j
				break
			}
			if intervals[j][0] <= intervals[i][1] {
				if intervals[j][1] <= intervals[i][1] {
					i = j + 1
					break
				} else {
					right = intervals[j][1]
					i = j + 1
					break
				}
			}
		}
		ans = append(ans, []int{left, right})

	}
	return ans
}

func main() {
	tar := make([][]int, 0)
	tar = append(tar, []int{1, 4})
	tar = append(tar, []int{5, 6})
	s := merge(tar)
	fmt.Println(s)
}
