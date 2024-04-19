package main

import "fmt"

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(matrix))
	}
	fmt.Println(res)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Println(j, i)
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
