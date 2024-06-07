package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for n := j + 1; n < len(nums); n++ {
				for m := n + 1; m < len(nums); m++ {
					if nums[i]+nums[j]+nums[n]+nums[m] == target {
						res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
					}
				}
			}
		}
	}
	return res
}
