package main

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	// 维护一个前缀乘积的slice
	pre := make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] * nums[i]
	}
	// 维护一个后缀乘积的slice
	post := make([]int, len(nums))
	post[len(post)-1] = nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		post[i] = post[i+1] * nums[i]
	}

	// 遍历
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			answer[i] = post[i+1]
			continue
		}
		if i == len(nums)-1 {
			answer[i] = pre[i-1]
			continue
		}
		answer[i] = pre[i-1] * post[i+1]

	}
	return answer
}
