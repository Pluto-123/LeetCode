package main

import "fmt"

// 示例 1:
//
// 输入: nums = [2,3,1,1,4]
// 输出: 2
// 解释: 跳到最后一个位置的最小跳跃数是 2。
//
//	从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
func jump(nums []int) int {
	steps := 0
	start := 0
	for {
		start = getNextIndex(nums, start)
		if start == -1 {
			// -1 表示已经在最终位置了
			return steps
		} else {
			steps++
			continue
		}
	}
}

func getNextIndex(nums []int, start int) int {
	// 判断是不是终点
	if start == len(nums)-1 {
		return -1
	}
	// 判断下一步能不能直接跳转到终点
	if nums[start]+start >= len(nums)-1 {
		return len(nums) - 1
	}
	// 贪心找到当前的最优解
	res := start
	maxReachedIndex := start
	for i := 1; i <= nums[start]; i++ {
		if nums[i+start]+i+start > maxReachedIndex {
			maxReachedIndex = nums[i+start] + i + start
			res = start + i
		}
	}

	return res
}

func main() {
	s := jump([]int{1, 2, 1, 1, 1})
	fmt.Println(s)
}
