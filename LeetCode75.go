package main

import "fmt"

func sortColors(nums []int) {
	quickSort75(nums, 0, len(nums)-1)
}

func quickSort75(nums []int, low int, high int) {
	if low > high {
		return
	}
	partition := partition75(nums, low, high)
	quickSort75(nums, low, partition-1)
	quickSort75(nums, partition+1, high)
}

func partition75(nums []int, low int, high int) int {
	left := low
	right := high
	for left < right {
		for left < right && nums[right] >= nums[low] {
			right--
		}
		for left < right && nums[left] <= nums[low] {
			left++
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[low] = nums[low], nums[left]
	return left
}

func main() {
	s := []int{0, 1, 2}
	sortColors(s)
	fmt.Println(s)
}
