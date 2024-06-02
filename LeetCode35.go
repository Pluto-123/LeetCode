package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
			continue
		}
		if nums[mid] < target {
			low = mid + 1
			continue
		}
	}
	return low
}
