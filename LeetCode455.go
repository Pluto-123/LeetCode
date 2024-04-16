package main

import "sort"

// success
func findContentChildren(g []int, s []int) int {
	// 尽量让胃口小的孩子吃小块的饼干
	res := 0
	sort.Ints(g)
	sort.Ints(s)
	//ig := 0
	is := 0
	for ig := 0; ig < len(g); ig++ {
		if is == len(s) {
			return res
		}
		if g[ig] <= s[is] {
			is++
			res++
		} else {
			is++
			// 对冲 ig++
			ig--
		}
	}
	return res
}

func main() {
	findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8})
}
