func solve(board [][]byte) {
    rows, cols := len(board), len(board[0])

	safe := make([][]bool, len(board))
	for row := range safe {
		safe[row] = make([]bool, len(board[0]))
	}

	var backtrack func(int, int)
	backtrack = func(r, c int) {
		if r == rows || c == cols || r < 0 || c < 0 || 
		board[r][c] == 'X' || safe[r][c] {return}

		safe[r][c] = true
		backtrack(r + 1, c)
		backtrack(r - 1, c)
		backtrack(r, c + 1)
		backtrack(r, c - 1)
	}

	for r := 0; r < rows; r++ {
		c := 0
		if board[r][c] == 'O' && !safe[r][c] {
			backtrack(r, c)
		}
		c = cols-1
		if board[r][c] == 'O' && !safe[r][c] {
			backtrack(r, c)
		}
	}

	for c := 0; c < cols; c++ {
		r := 0
		if board[r][c] == 'O' && !safe[r][c] {
			backtrack(r, c)
		}
		r = rows-1
		if board[r][c] == 'O' && !safe[r][c] {
			backtrack(r, c)
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !safe[r][c] {board[r][c] = 'X'}
		}
	}
}