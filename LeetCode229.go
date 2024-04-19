package main

func majorityElement(nums []int) []int {
	times := len(nums) / 3

	res := make([]int, 0)
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := mp[nums[i]]
		if ok {
			mp[nums[i]]++
		} else {
			mp[nums[i]] = 1
		}
	}
	for key, value := range mp {
		if value > times {
			res = append(res, key)
		} else {
			continue
		}
	}
	return res
}
