func isValidSudoku(board [][]byte) bool {
	seenRows := make([]int, 9)	// if bit 1 is 1: number "1" appeared in row
	seenCols := make([]int, 9)
	seenBoxes := make([]int, 9)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				continue
			}
			num := int(board[row][col] - '1')	// minus 1 to prevent overflow
			mask := 1 << num
			box := (row/3)*3+col/3
			if (seenRows[row] & mask != 0) || (seenCols[col] & mask != 0) || (seenBoxes[box] & mask != 0) {
				return false
			}
			seenRows[row] |= mask
			seenCols[col] |= mask
			seenBoxes[box] |= mask
		}
	}

	return true
}
