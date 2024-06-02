package main

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		if target >= matrix[i][0] && target <= matrix[i][n-1] {
			return binarySearch74(matrix[i], target)
		}
	}
	return false
}

func binarySearch74(nums []int, target int) bool {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] < target {
			low = mid + 1
		}
	}
	return false

}
