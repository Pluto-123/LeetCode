package main

import "fmt"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	// 初始化
	for i := 1; i < len(dp); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			dp[i][j] = grid[i][j] + min(dp[i][j-1], dp[i-1][j])
		}
	}
	fmt.Println(dp)
	return dp[len(dp)-1][len(dp[0])-1]
}
