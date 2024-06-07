package main

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	pre := make([]int, len(nums))
	post := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			pre[i] = nums[i]
			continue
		}
		pre[i] = pre[i-1] * nums[i]
	}

	for i := len(post) - 1; i >= 0; i-- {
		if i == len(post)-1 {
			post[i] = nums[i]
			continue
		}
		post[i] = nums[i] * post[i+1]
	}

	for i := 0; i < len(res); i++ {
		if i == 0 {
			res[i] = post[i+1]
			continue
		}
		if i == len(res)-1 {
			res[i] = pre[i-1]
			continue
		}
		res[i] = post[i+1] * pre[i-1]
	}

	return res
}
