package main

func canPartition(nums []int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	if len(nums) < 2 {
		return false
	}
	dp := make([][]bool, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum/2+1)
	}

	// 初始化第一列
	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	// 初始化第一行
	dp[0][nums[0]] = true

	for i := 1; i < len(dp); i++ {
		v := nums[i]
		for j := 1; j <= sum/2; j++ {
			if j < v {
				dp[i][j] = dp[i-1][j]
			}
			if j == v {
				dp[i][j] = true
			}
			if j > v {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			}
		}
	}
	return dp[len(nums)-1][sum/2]

}
