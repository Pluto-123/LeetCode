package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	//dp[0] = 0
	for i := 1; i < len(dp); i++ {
		if coins[i] > amount {
			dp[i] = dp[i-1]
		}
		if coins[i] == amount {
			dp[i] = 1
		}
		if coins[i] < amount {
			dp[i] = min(dp[i-1], dp[amount/coins[i]]+dp[amount%coins[i]])
		}
	}
	return dp[amount]
}
func main() {
	coinChange([]int{1, 2, 5}, 11)
}
