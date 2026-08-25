func numIslands(grid [][]byte) int {
	directions := [][2]int{{1,0},{-1,0},{0,1},{0,-1}}
	rows, cols := len(grid), len(grid[0])
    res := 0

	var bfs func(r, c int)
	bfs = func(r, c int) {
		q := [][2]int{{r,c}}
		grid[r][c] = '0'

		for len(q) > 0 {
			row, col := q[0][0], q[0][1]
			q = q[1:]

			for _, d := range directions {
				dr, dc := d[0] + row, d[1] + col

				if dr < rows && dr > -1 && dc < cols && dc > -1 && grid[dr][dc] == '1' { 
					q = append(q, [2]int{dr, dc})
					grid[dr][dc] = '0'
				}
			}
		}

	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				bfs(r, c) 
				res++
			}
		}
	}

	return res
}