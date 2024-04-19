package main

import (
	"sort"
)

func matrixSum(nums [][]int) int {
	res := 0
	m := len(nums)
	n := len(nums[0])
	for i := 0; i < m; i++ {
		sort.Ints(nums[i])
	}
	for j := 0; j < n; j++ {
		maxVal := 0
		for i := 0; i < m; i++ {
			if nums[i][j] > maxVal {
				maxVal = nums[i][j]
			}
		}
		res += maxVal
	}
	return res
}
