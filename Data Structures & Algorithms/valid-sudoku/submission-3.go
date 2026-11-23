func isValidSudoku(board [][]byte) bool {
	seenRows := [9][9]bool{}	// row => num
	seenCols := [9][9]bool{}	// col => num
	seenBoxes := [9][9]bool{}	// box => num
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			sVal := string(board[row][col])
			if sVal == "." {
				continue
			}
			num, _ := strconv.Atoi(sVal)
			num--	// prevent over flow
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
