package main

func rob(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}

	ans := 0
	// dp[i]表示偷第i个的情况下最多偷多少
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[0] + nums[2]
	ans = max(max(max(ans, dp[1]), dp[0]), dp[2])
	for i := 3; i < len(dp); i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
		ans = max(ans, dp[i])
	}
	return ans
}
