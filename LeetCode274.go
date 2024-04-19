package main

func hIndex(citations []int) int {
	//h := 0
	d := make([]int, len(citations)+1)

	for i := 0; i < len(citations); i++ {
		cites := citations[i]
		if cites >= 0 && cites < len(d) {
			for j := 0; j <= cites; j++ {
				d[j]++
			}
		} else {
			// 全部都加1
			for j := 0; j < len(d); j++ {
				d[j]++
			}
		}
	}
	for i := len(d) - 1; i >= 0; i-- {
		if d[i] >= i {
			return i
		}
	}

	return -1
}
