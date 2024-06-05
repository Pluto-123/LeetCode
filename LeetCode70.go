package main

func climbStairs(n int) int {
	dp := make([]int, n+2)
	dp[1] = 1
	dp[0] = 1

	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
