package main

import "fmt"

func findLHS(nums []int) int {
	ans := 0
	//left := 0
	//right := 0
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		_, ok := mp[nums[i]]
		if ok {
			mp[nums[i]]++
		} else {
			mp[nums[i]] = 1
		}
	}
	for key, value := range mp {
		fmt.Println(key, value)
		ots := max(mp[key+1], mp[key-1])
		fmt.Println(ots)
		if ots == 0 {
			continue
		}
		ans = max(ans, value+ots)
	}
	return ans
}
