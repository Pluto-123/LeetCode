package main

import "fmt"

// 给你一个仅由整数组成的有序数组，其中每个元素都会出现两次，唯有一个数只会出现一次。
//
// 请你找出并返回只出现一次的那个数。
//
// 你设计的解决方案必须满足 O(log n) 时间复杂度和 O(1) 空间复杂度。
func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	low := 0
	high := len(nums) - 1
	res := bsearch(nums, low, high)
	return res
}

// -1
func bsearch(nums []int, low int, high int) int {
	for low <= high {
		mid := (low + high) / 2
		if success(nums, mid) {
			return nums[mid]
		}
		tmp := bsearch(nums, low, mid-1)
		if tmp != -1 {
			return tmp
		}
		return bsearch(nums, mid+1, high)
	}
	return -1
}

func success(nums []int, mid int) bool {
	if mid == 0 {
		return !(nums[mid] == nums[mid+1])
	}
	if mid == len(nums)-1 {
		return !(nums[mid] == nums[mid-1])
	}
	return nums[mid] != nums[mid-1] && nums[mid] != nums[mid+1]
}

//[1,1,2,3,3,4,4,8,8]

func main() {
	s := singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
	fmt.Println(s)
}
