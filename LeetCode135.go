package main

func candy(ratings []int) int {
	ans := 0

	return ans
}

func candyDG(ratings []int, sum int) int {
	ans := 0
	ans += 1
	dp := make([]int, 0)
	dp = append(dp, 1)
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			//dp = append(dp, dp[len(dp)-1]+1)
		}
		if ratings[i] < ratings[i-1] {
			dp = append(dp, dp[len(dp)-1]-1)
		}
		if ratings[i] == ratings[i-1] {

			break
		}
	}
	return ans
}
