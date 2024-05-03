package main

import "sort"

func average(salary []int) float64 {
	sort.Ints(salary)
	res := 0
	for i := 1; i < len(salary)-1; i++ {
		res += salary[i]
	}
	return float64(res) / float64(len(salary)-2)
}
