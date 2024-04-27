package main

import "fmt"

func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	left := 0
	right := 0
	dp := make([][]int, 0)

	for {
		// 跳出判断 有问题
		if right == len(nums) {
			break
		}
		m := mul(nums, left, right)
		if m < k {
			right++
		}
		// 让right恒大于等于left
		if m >= k {
			// 存储
			dp = append(dp, []int{left, right})
			left = right + 1
			right = left
		}
	}
	// 做dp的阶乘
	for i := 0; i < len(dp); i++ {
		ans += jiecheng(dp[i][1] - dp[i][0] + 1)
	}
	return ans

}

func jiecheng(i int) int {
	res := 1
	for k := 1; k <= i; k++ {
		res *= k
	}
	return res
}

func mul(nums []int, left int, right int) int {
	res := 1
	for i := left; i <= right; i++ {
		res *= nums[i]
	}
	return res
}

func main() {
	s := numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100)
	fmt.Println(s)
}
