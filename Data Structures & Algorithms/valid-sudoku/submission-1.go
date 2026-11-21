func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9][9]bool
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			str := string(board[row][col])
			if str == "." {
				continue
			}
			val, _ := strconv.Atoi(str)
			val--
			if rows[row][val] || cols[col][val] || boxes[(row/3)*3+col/3][val] {
				return false
			}
			rows[row][val] = true
			cols[col][val] = true
			boxes[(row/3)*3+col/3][val] = true
		}
	}

	return true
}
