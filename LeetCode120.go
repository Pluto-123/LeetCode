package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	//dp := make([][]int, len(triangle))
	//for i := 0; i < len(dp); i++ {
	//	dp[i] = make([]int, len(triangle[i]))
	//}
	//
	//copy(dp, triangle)
	//for i := 0; i < len(dp); i++ {
	//	for j := 0; j < len(dp[i]); j++ {
	//		dp[i][j] = 0
	//	}
	//}
	dp := make([][]int, len(triangle))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(triangle[i]))
		copy(dp[i], triangle[i]) // 复制triangle中的值到新的切片中
	}
	dp[0][0] = triangle[0][0]
	// 更新dp
	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(dp[i]); j++ {
			// 如果是左顶点
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
				continue
			}
			// 如果是最后一位
			if j == len(dp[i-1]) {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
				continue
			}
			// 如果是中间的节点
			if j+1 == len(dp[i-1]) {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
		}
	}
	//dp[0][0] = 1
	fmt.Println(dp)
	res := dp[len(dp)-1][len(dp[len(dp)-1])-1]
	for i := 0; i < len(dp[len(dp)-1]); i++ {
		res = min(res, dp[len(dp)-1][i])
	}
	return res
}
