package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// 三指针
	res := make([][]int, 0)
	p := 0
	q := len(nums) - 1
	k := p + 1
	for p < q {
		if k == p || k == q {
			p++
			k = p + 1
			continue
		}
		if nums[p]+nums[k]+nums[q] == 0 {
			res = append(res, []int{nums[p], nums[k], nums[q]})
			p++
			q--
			k = p + 1
		}
		if nums[p]+nums[k]+nums[q] != 0 {
			k++
		}

	}
	return quchong(res)
}

func quchong(target [][]int) [][]int {
	for i := 0; i < len(target); i++ {
		sort.Ints(target[i])
	}
	return removeDuplicates2(target)

}

func removeDuplicates2(slices [][]int) [][]int {
	// 使用 map 存储已经出现过的切片
	uniqueSlices := make(map[string]bool)
	var result [][]int

	for _, slice := range slices {
		// 将切片转换为字符串，作为 map 的键
		sliceStr := fmt.Sprintf("%v", slice)

		// 如果切片不在 map 中，则添加到结果中并标记为已出现
		if !uniqueSlices[sliceStr] {
			uniqueSlices[sliceStr] = true
			result = append(result, slice)
		}
	}

	return result
}

func main() {
	s := threeSum([]int{1, -1, -1, 0})
	fmt.Println(s)
}
