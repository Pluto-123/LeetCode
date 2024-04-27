package main

func sortArray(nums []int) []int {
	quickSort3(nums, 0, len(nums)-1)
	return nums
}

func quickSort3(nums []int, low int, high int) {
	if low >= high {
		return
	}
	partition := partition31(nums, low, high)
	quickSort3(nums, low, partition-1)
	quickSort3(nums, partition+1, high)
}

func partition31(nums []int, left int, right int) int {
	i := left
	j := right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}
