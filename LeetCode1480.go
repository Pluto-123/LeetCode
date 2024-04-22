package main

func runningSum(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] + nums[i]
	}
	return res
}
