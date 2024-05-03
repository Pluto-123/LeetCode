package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0)
	mp := map[int]int{}
	d := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		_, ok := mp[nums[i]]
		if ok {
			mp[nums[i]]++
		} else {
			mp[nums[i]] = 1
		}
	}
	for key, value := range mp {
		d = append(d, []int{key, value})
	}
	quickSort4(d, 0, len(d)-1)
	fmt.Println(d)
	for i := 0; i < k; i++ {
		res = append(res, d[(len(d) - 1)][0])
		d = d[:len(d)-1]
	}
	return res
}

func quickSort4(d [][]int, low int, high int) {
	if low > high {
		return
	}
	partition := partition4(d, low, high)
	quickSort4(d, low, partition-1)
	quickSort4(d, partition+1, high)
}

func partition4(d [][]int, low int, high int) int {
	i := low
	j := high
	for i < j {
		for i < j && d[j][1] >= d[low][1] {
			j--
		}
		for i < j && d[i][1] <= d[low][1] {
			i++
		}
		d[i], d[j] = d[j], d[i]
	}
	d[i], d[low] = d[low], d[i]
	return i

}
