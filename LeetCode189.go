package main

// 轮转数组
func rotate1(nums []int, k int) {
	//for i := 0; i < k; i++ {
	//	r189(nums)
	//}
	for k > len(nums) {
		k -= len(nums)
	}

	nums_2 := make([]int, 0)
	for i := len(nums) - k; i < len(nums); i++ {
		nums_2 = append(nums_2, nums[i])
	}
	for i := 0; i < len(nums)-k; i++ {
		nums_2 = append(nums_2, nums[i])
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = nums_2[i]
	}

}
