package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	path := make([]int, 0)
	backTrace3(nums, &path, &res)
	return res
}

func backTrace3(nums []int, path *[]int, res *[][]int) {
	if len(*path) == 1 {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(nums); i++ {

	}
}
