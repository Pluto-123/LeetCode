package main

import "math"

func maxSubArray(nums []int) int {
	ans := (-1) * math.MaxInt32
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
	}
	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}
