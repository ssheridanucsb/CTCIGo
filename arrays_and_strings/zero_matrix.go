package arraysandstrings

func zeroMatrix(matrix [][]int, M int, N int) {
	//finds zeros elements then zeros their corresponding rows and columns
	// M = number of rows
	// N = number of cols
	//keep track of seen indexes without introducing duplicates
	rowIdx := make(map[int]int)
	colIdx := make(map[int]int)

	//find zero indexes
	for i := 0; i < M; i++ {
		for j := 0; j < M; j++ {
			if matrix[i][j] == 0 {
				rowIdx[i]++
				colIdx[j]++
			}
		}
	}

	//iterate over mapped indexes and zero out the correspondig row/col
	for key := range rowIdx {
		for i := 0; i < N; i++ {
			matrix[key][i] = 0
		}
	}

	for key := range colIdx {
		for i := 0; i < N; i++ {
			matrix[i][key] = 0
		}
	}

}
