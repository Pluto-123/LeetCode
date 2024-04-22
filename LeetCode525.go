package main

import "fmt"

func findMaxLength(nums []int) int {
	dp := make([]int, len(nums))
	res := 0
	mp := map[int]int{}
	if nums[0] == 0 {
		dp[0] = -1
	} else {
		dp[0] = 1
	}
	mp[dp[0]] = 0

	for i := 1; i < len(dp); i++ {
		if nums[i] == 0 {
			dp[i] = dp[i-1] - 1
		}
		if nums[i] == 1 {
			dp[i] = dp[i-1] + 1
		}
		// 以dp[i]为key， index为value
		_, ok := mp[dp[i]]
		if ok {
			res = max(res, i-mp[dp[i]])
		}
		if dp[i] == 0 {
			res = max(res, i+1)
		}

		if !ok && dp[i] == 0 {
			res = max(res, i+1)
			mp[dp[i]] = i
			continue
		}

		if !ok {
			mp[dp[i]] = i
			continue
		}

	}

	return res
}

func main() {
	s := findMaxLength([]int{0, 1, 0, 1})
	fmt.Println(s)
}
