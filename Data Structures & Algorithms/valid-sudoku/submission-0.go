func isValidSudoku(board [][]byte) bool {
	// Check column
	for i := 0; i < 9; i++ {
		seen := make(map[int]struct{})
		for j := 0; j < 9; j++ {
			str := string(board[i][j])
			if str != "." {
				key, _ := strconv.Atoi(str)
				if _, ok := seen[key]; ok {
					return false
				}
				seen[key] = struct{}{}
			}
		}
	}

	// Check row
	for i := 0; i < 9; i++ {
		seen := make(map[int]struct{})
		for j := 0; j < 9; j++ {
			str := string(board[j][i])
			if str != "." {
				key, _ := strconv.Atoi(str)
				if _, ok := seen[key]; ok {
					return false
				}
				seen[key] = struct{}{}
			}
		}
	}

	// Check sub-boxes
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			seen := make(map[int]struct{})
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					str := string(board[3*i+k][3*j+l])
					if str != "." {
						key, _ := strconv.Atoi(str)
						if _, ok := seen[key]; ok {
							return false
						}
						seen[key] = struct{}{}
					}
				}
			}
		}
	}

	return true
}
