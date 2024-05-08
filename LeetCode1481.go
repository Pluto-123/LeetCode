package main

import "fmt"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	mp := map[int]int{}
	for i := 0; i < len(arr); i++ {
		_, ok := mp[arr[i]]
		if ok {
			mp[arr[i]]++
		} else {
			mp[arr[i]] = 1
		}
	}

	s := make([][]int, 0)
	for key, value := range mp {
		s = append(s, []int{key, value})
	}
	quickSort1481(s, 0, len(s)-1)
	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		k -= s[i][1]
		if k > 0 {
			continue
		}
		if k == 0 {
			return len(s) - i - 1
		}
		if k < 0 {
			return len(s) - i
		}

	}
	return 0
}

func quickSort1481(s [][]int, low int, high int) {
	if low > high {
		return
	}
	partition := partition1481(s, low, high)
	quickSort1481(s, low, partition-1)
	quickSort1481(s, partition+1, high)
}

func partition1481(s [][]int, low int, high int) int {
	i := low
	j := high
	for i < j {
		for i < j && s[j][1] >= s[low][1] {
			j--
		}
		for i < j && s[i][1] <= s[low][1] {
			i++
		}
		s[i], s[j] = s[j], s[i]
	}
	s[i], s[low] = s[low], s[i]
	return i
}
