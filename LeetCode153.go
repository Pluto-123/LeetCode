package main

import "fmt"

func findMin(nums []int) int {
	//ans := 0
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2

		if nums[mid] > nums[left] && nums[mid] < nums[right] {
			// 左右都有序
			return nums[left]
		}
		if nums[mid] > nums[left] {
			// 左有序
			left = mid + 1
			continue
		}
		if nums[mid] < nums[right] {
			// 右有序
			right = mid - 1
			continue
		}
		if nums[mid] == nums[left] {
			return min(nums[mid], nums[right])
		}

	}
	return nums[left]
}

func main() {
	s := findMin([]int{3, 1, 2})
	fmt.Println(s)
}
