package main

func firstMissingPositive(nums []int) int {
	mp := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if mp[nums[i]] {
			mp[nums[i]-1] = false
		} else {
			mp[nums[i]] = true
			mp[nums[i]-1] = false
		}
	}

	for k, v := range mp {
		if v == false {
			return k
		}
	}
	return 0
}
