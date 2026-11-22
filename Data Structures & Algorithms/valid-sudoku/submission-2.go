func isValidSudoku(board [][]byte) bool {
	seenRow := [9][9]bool{}	// seen number in row[int]
	seenCol := [9][9]bool{}	// seen number in col[int]
	seenBox := [9][9]bool{}	// seen number in box[int] (ex: box num in row 0 is 0,1,2)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			num, _ := strconv.Atoi(string(board[row][col]))
			if seenRow[row][num-1] || seenCol[col][num-1] || seenBox[(row/3)*3+col/3][num-1] {
				return false
			}
			seenRow[row][num-1] = true
			seenCol[col][num-1] = true
			seenBox[(row/3)*3+col/3][num-1] = true
		}
	}

	return true
}
