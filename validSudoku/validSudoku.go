package validSudoku

// ValidSudoku check if sudoku is valid row wise column wise, 3*3 matrix wise
func ValidSudoku(board [][]byte) bool {

	// check line
	for _, col := range board {
		Map := map[byte]int{}
		for _, e := range col {
			if e != '.' {
				if _, ok := Map[e]; ok {
					return false
				} else {
					Map[e] = 1
				}
			}
		}
	}
	// check column

	for i, col := range board {
		Map := map[byte]int{}
		for j := range col {
			e := board[j][i]
			if e != '.' {
				if _, ok := Map[e]; ok {
					return false
				} else {
					Map[e] = 1
				}
			}
		}
	}
	// check 3x3 box
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isSame3X3(board, i, j) {
				return false
			}
		}
	}
	return true
}

// check 3X3 box
func isSame3X3(board [][]byte, m, n int) bool {
	Map := map[byte]int{}
	for i := 3 * m; i < 3*(m+1); i++ {
		for j := 3 * n; j < 3*(n+1); j++ {
			e := board[i][j]
			if e != '.' {
				// fmt.Println(e, i, j)
				if _, ok := Map[e]; ok {
					return false
				} else {
					Map[e] = 1
				}
			}
		}
	}
	return true
}
