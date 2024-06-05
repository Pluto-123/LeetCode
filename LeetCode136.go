package main

func singleNumber2(nums []int) int {
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		_, ok := mp[nums[i]]
		if ok {
			delete(mp, nums[i])
		} else {
			mp[nums[i]] = false
		}
	}
	for key, _ := range mp {
		return key
	}
	return 0
}

// 位运算的解法
func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}
