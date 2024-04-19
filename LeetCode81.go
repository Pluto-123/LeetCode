package main

import "sort"

func search2(nums []int, target int) bool {
	sort.Ints(nums)
	return bS(nums, target)
}

func bS(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] < target {
			low = mid + 1
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] == target {
			return true
		}
	}
	return false
}
