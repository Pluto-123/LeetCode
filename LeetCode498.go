package main

func findDiagonalOrder(mat [][]int) []int {
	res := make([]int, 0)

	for k := 0; k < len(mat); k++ {
		temp := []int{}

		for i := 0; i <= k; i++ {
			j := k - i
			temp = append(temp, mat[i][j])
		}

		if k%2 == 1 {
			reverseSlicce(temp)
		}
		res = append(res, temp...)
	}
	return res
}

func reverseSlicce(temp []int) {
	left := 0
	right := len(temp) - 1

	for left < right {
		temp[left], temp[right] = temp[right], temp[left]
		left++
		right--
	}
}
