package main

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		if bSearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func bSearch(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			//left = mid + 1
			right = mid - 1
		}
		if nums[mid] < target {
			left = mid + 1
		}
	}
	return false
}
