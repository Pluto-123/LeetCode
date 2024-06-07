package main

func rotate(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i >= j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}
		} // for
	} //for

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[0])-1-j] = matrix[i][len(matrix[0])-1-j], matrix[i][j]
		}
	}
	//for i := 0; i < len(matrix[0])/2; i++ {
	//	matrix[:][i], matrix[:][len(matrix)-i] = matrix[:][len(matrix)-i], matrix[:][i]
	//}
}
