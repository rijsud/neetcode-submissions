func islandsAndTreasure(grid [][]int) {
    INF := 2147483647
    rows, cols := len(grid), len(grid[0])
    dirs := [][]int{{0,1},{0,-1},{1,0},{-1,0}}
    q := [][2]int{}
    
    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 0 {q = append(q, [2]int{r,c})}
        }
    }

    for len(q) > 0 {
        row, col := q[0][0], q[0][1]
        q = q[1:]
        for _, dir := range dirs {
            nr, nc := row + dir[0], col + dir[1]
            if nr == rows || nr < 0 || nc == cols || nc < 0 ||
            grid[nr][nc] != INF {continue}
            q = append(q, [2]int{nr,nc})
            grid[nr][nc] = grid[row][col] + 1
        }
    }
}
