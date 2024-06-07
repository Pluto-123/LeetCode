package main

func setZeroes(matrix [][]int) {
	todo := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				todo = append(todo, []int{i, j})
			}
		} // for
	} // for
	for i := 0; i < len(todo); i++ {

		setZero(matrix, todo[i][0], todo[i][1])
	}
}

func setZero(matrix [][]int, x int, y int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][y] = 0
	}
	for i := 0; i < len(matrix[0]); i++ {
		matrix[x][i] = 0
	}
}
