package main

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)
	}
	// 初始化边界
	for i := 0; i < numRows; i++ {
		dp[i][0] = 1
		dp[i][len(dp[i])-1] = 1
	}
	for i := 2; i < numRows; i++ {
		for j := 1; j <= len(dp[i])-2; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}
	return dp
}
