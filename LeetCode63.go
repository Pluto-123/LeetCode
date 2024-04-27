package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	if obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1] == 1 {
		return 0
	}
	if len(obstacleGrid) == 1 {
		return 1
	}
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(obstacleGrid[0]))
	}
	dp[0][0] = 1
	// 先初始化第一列
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}

	// 初始化第一行
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0][i] = 0
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}
	fmt.Println(dp)
	// 更新中间的dp
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	//fmt.Println(dp)
	return dp[len(dp)-1][len(dp[0])-1]
}
