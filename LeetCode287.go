package main

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return i
			}
		}
	}
	return 0
}
