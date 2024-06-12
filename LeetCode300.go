package main

import "fmt"

func lengthOfLIS(nums []int) int {
	ans := 1
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		if find300(nums, dp, i) == -1 {
			dp[i] = 1
		} else {
			dp[i] = dp[find300(nums, dp, i)] + 1
		}
	}

	for i := 0; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	fmt.Println(dp)
	return ans
}

func find300(nums []int, dp []int, start int) int {

	for i := start - 1; i >= 0; i-- {
		if nums[i] < nums[start] {
			return i
		}
	}
	return -1
}
