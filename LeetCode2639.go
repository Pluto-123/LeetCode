package main

func findColumnWidth(grid [][]int) []int {
	res := make([]int, 0)
	for j := 0; j < len(grid[0]); j++ {
		// 列
		m := 0
		for i := 0; i < len(grid); i++ {
			m = max(m, getLength(grid[i][j]))
		}
		res = append(res, m)
	}
	return res
}

func getLength(num int) int {
	res := 0
	if num < 0 {
		num = (-1) * num
		res++
	}
	for num >= 10 {
		num = num / 10
		res++
	}
	res++
	return res
}

func main() {
	t := make([][]int, 0)
	t = append(t, []int{1})
	findColumnWidth(t)
}
