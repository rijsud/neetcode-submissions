func orangesRotting(grid [][]int) int {
    res, fresh := 0, 0
    rows, cols := len(grid), len(grid[0])
	directions := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}
	q := [][2]int{}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {q = append(q, [2]int{r,c})}
			if grid[r][c] == 1 {fresh++}
		}
	}

	for len(q) > 0 && fresh > 0 {
		qLen := len(q)
		for i := 0; i < qLen; i++ {
			row, col := q[0][0], q[0][1]
			q = q[1:]

			for _, dir := range directions {
				r, c := row + dir[0], col + dir[1]
				if r == rows || c == cols || r < 0 || c < 0 ||
				grid[r][c] != 1 {continue}
				grid[r][c] = 2
				fresh--
				q = append(q, [2]int{r,c})
			}
		}
		res++
	}

	if fresh == 0 {return res} else {return -1}
}
