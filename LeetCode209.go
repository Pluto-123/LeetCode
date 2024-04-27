package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口

	ans := math.MaxInt32
	left := 0
	right := 0

	sum := nums[0]

	//for {
	//	if right == len(nums)-1 && sum < target {
	//		break
	//	}
	//	if sum >= target && left < len(nums) {
	//		ans = min(ans, right-left+1)
	//		sum -= nums[left]
	//		left++
	//	}
	//	if sum < target && right < len(nums) {
	//		right++
	//		sum += nums[right]
	//	}
	//}
	for right < len(nums) {
		if sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left++
		}
		if sum < target {
			right++
			if right == len(nums) {
				continue
			}
			sum += nums[right]
			//ans = min(ans, right-left+1)
		}
	}

	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
