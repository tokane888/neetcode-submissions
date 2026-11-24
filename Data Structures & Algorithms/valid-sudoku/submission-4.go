func isValidSudoku(board [][]byte) bool {
	seenRows := [9][9]bool{}	// row => num
	seenCols := [9][9]bool{}	// col => num
	seenBoxes := [9][9]bool{}	// box => num
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			num, _ := strconv.Atoi(string(board[row][col]))
			num--	// prevent overflow
			if seenRows[row][num] || seenCols[col][num] || seenBoxes[(row/3)*3+col/3][num] {
				return false
			}
			seenRows[row][num] = true
			seenCols[col][num] = true
			seenBoxes[(row/3)*3+col/3][num] = true
		}
	}

	return true
}
