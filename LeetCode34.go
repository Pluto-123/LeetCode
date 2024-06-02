package main

func searchRange(nums []int, target int) []int {
	//low := 0
	//high := len(nums) - 1
	left := 0
	right := len(nums) - 1
	if binarySearch34(nums, 0, len(nums)-1, target) == -1 {
		return []int{-1, -1}
	}

	//for binarySearch34(nums, left, right, target) != -1 {
	//	right = binarySearch34(nums, left, right, target)
	//}
	//left = 0
	//right = len(nums) - 1
	//for binarySearch34(nums, left, right, target) != -1 {
	//	left = binarySearch34(nums, left, right, target)
	//}

	return []int{right, left}
}

func binarySearch34(nums []int, low, high, target int) int {
	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] < target {
			low = mid + 1
		}
	}
	return -1
}
