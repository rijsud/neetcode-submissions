func solveNQueens(n int) [][]string {
	cols := map[int]bool{}
	posDiag := map[int]bool{} // r + c
	negDiag := map[int]bool{} // r - c

	res := [][]string{}
	board := make([][]string, n)

	for i := range board {
		board[i] = make([]string, n)
		for j := range board[i] {
			board[i][j] = "."
		}
	}

	var backtrack func(int)
	backtrack = func(r int) {
		if r == n {
			boardCopy := make([]string, n)
			for i := range board{
				boardCopy[i] = ""
				for j := range board[i] {
					boardCopy[i] += board[i][j]
				}
			}
			res = append(res, boardCopy)
			return
		}

		for c := 0; c < n; c++ {
			if cols[c] || posDiag[r+c] || negDiag[r-c] { continue }

			cols[c], posDiag[r+c], negDiag[r-c] = true, true, true
			board[r][c] = "Q"

			backtrack(r+1)

			cols[c], posDiag[r+c], negDiag[r-c] = false, false, false
			board[r][c] = "."
		}
	}

	backtrack(0)
	return res
}