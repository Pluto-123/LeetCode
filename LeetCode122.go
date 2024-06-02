package main

// 第一次从零推导出来的dp+贪心的状态转移方程
// 用最少的步数得到最多的利润
func maxProfit2(prices []int) int {

	if len(prices) < 2 {
		return 0
	}
	// dpi表示以i为结尾卖出的最大利润
	dp := make([]int, len(prices))
	dp[0] = 0
	// 初始化前两个值完成
	if prices[1] >= prices[0] {
		dp[1] = prices[1] - prices[0]
	} else {
		dp[1] = 0
	}

	for i := 2; i < len(dp); i++ {
		tempMax := 0 // 用来暂存最大值
		if prices[i] <= prices[i-1] {
			dp[i] = dp[i-1]
		} else {
			// 如果新加入进来的price大于前一天的price
			for j := 0; j < i; j++ {
				if j == 0 {
					tempMax = max(tempMax, prices[i]-prices[j])
				} else {
					tempMax = max(tempMax, prices[i]-prices[j]+dp[j-1])
				}
			}
			dp[i] = tempMax
		}

	}
	return dp[len(dp)-1]
}
