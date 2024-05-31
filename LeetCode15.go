package main

import (
	"sort"
	"strconv"
	"strings"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	for k := 0; k < len(nums)-2; k++ {
		if nums[k] > 0 {
			break
		}
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		i := k + 1
		j := len(nums) - 1
		for {
			if i >= j {
				break
			}
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				i++
			}
			if nums[i]+nums[j]+nums[k] > 0 {
				j--
			}
			if nums[i]+nums[j]+nums[k] < 0 {
				i++
			}
		}

	}
	return quchong15(ans)
}

func quchong15(target [][]int) [][]int {
	res := make([][]int, 0)
	mp := make(map[string]bool)
	for i := 0; i < len(target); i++ {
		sb := strings.Builder{}
		for j := 0; j < len(target[i]); j++ {
			sb.WriteString(strconv.Itoa(target[i][j]))
		}
		if _, ok := mp[sb.String()]; ok {
			continue
		} else {
			mp[sb.String()] = true
			res = append(res, target[i])
		}
	}
	return res
}
