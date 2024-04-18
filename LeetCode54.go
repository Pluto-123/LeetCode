package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	// 0 1 2 3 分别标识右、下、左、上
	left := 0 // 闭区间
	right := len(matrix[0]) - 1
	top := 0
	bottom := len(matrix) - 1
	direction := make([]bool, 4)
	direction[0] = true

	res := make([]int, 0)
	// 如果数量不相同
	trace := make([]int, 2)
	for !(len(matrix)*len(matrix[0]) == len(res)) {
		res = append(res, matrix[trace[0]][trace[1]])
		if changeDirection(direction, left, right, top, bottom, trace) {
			//  changeD
			for i := 0; i < len(direction); i++ {
				if direction[i] == true {
					direction[i] = false
					if i == 0 {
						top++
					}
					if i == 1 {
						right--
					}
					if i == 2 {
						bottom--
					}
					if i == 3 {
						left++
					}

					direction[(i+1)%len(direction)] = true
					break
				}
			}

		}
		// 走一步
		if direction[0] {
			trace[1]++
			continue
		}
		if direction[1] {
			trace[0]++

			continue
		}
		if direction[2] {
			trace[1]--
			continue
		}
		if direction[3] {
			trace[0]--
			continue
		}

	}
	return res
}

func changeDirection(direction []bool, left int, right int, top int, bottom int, trace []int) bool {
	// 先把当前的方向拿出来
	d := 0
	for i := 0; i < len(direction); i++ {
		if direction[i] == true {
			d = i
			break
		}
	}
	// 右
	if d == 0 {
		if trace[1] == right {
			return true
		}
	}

	if d == 1 {
		if trace[0] == bottom {
			return true
		}
	}
	if d == 2 {
		if trace[1] == left {
			return true
		}
	}
	if d == 3 {
		if trace[0] == top {
			return true
		}
	}
	return false
}

func main() {
	m := make([][]int, 0)
	m = append(m, []int{1, 2, 3, 4})
	m = append(m, []int{5, 6, 7, 8})
	m = append(m, []int{9, 10, 11, 12})
	s := spiralOrder(m)
	fmt.Println(s)
}
