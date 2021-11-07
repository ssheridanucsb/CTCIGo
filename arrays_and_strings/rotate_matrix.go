package arraysandstrings

func rotateMatrix(matrix [][]int) {
	//rotate a matrix in place
	//takes a slide of [][]int so all assignments are reflected on the undelrying arrays
	// we'll rotate the perimiter of the matrix then move in until we reach a 2x2 or 3x3 matrix
	N := len(matrix[0])
	startIdx := 0

	//lets do our rotations using a temporary array for swapping
	tmp := make([]int, N)
	for {

		//lets copy the first row
		for i := startIdx; i < N; i++ {
			tmp[i] = matrix[startIdx][i]
		}

		//now lets move the first column to the first row, inserting backwards
		for i := startIdx; i < N; i++ {
			matrix[startIdx][N-i-1+startIdx] = matrix[i][startIdx]
		}

		//next lets move the last row to the first column
		for i := startIdx; i < N; i++ {
			matrix[i][startIdx] = matrix[N-1][i]
		}

		//then we'll move the last column with the last row, inserting backwards
		for i := startIdx; i < N; i++ {
			matrix[N-1][i] = matrix[N-1-i+startIdx][N-1]
		}

		//finally we'll move our copy of the first row to the last column
		for i := startIdx; i < N; i++ {
			matrix[i][N-1] = tmp[i]
		}

		//if the matrix is 2x2 or 3x3 we're done, otherwise we need to do this again on the submatrix
		if N == 2 || N == 3 {
			return
		}

		//increment to rotate the sub matrix
		startIdx++
		N = N - 1
	}

}
