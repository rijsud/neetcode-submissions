func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
    q := [][2]int{}

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 0 {
                q = append(q, [2]int{r, c})
            }
        }
    }

    dirs := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

    for len(q) > 0 {
        row, col := q[0][0], q[0][1]
        dist := grid[row][col] + 1
        q = q[1:]

        for _, dir := range dirs {
            r, c := row + dir[0], col + dir[1]
            if r == rows || c == cols || r < 0 || c < 0 ||
               grid[r][c] != 2147483647 {
                continue
            }
            q = append(q, [2]int{r, c})
            grid[r][c] = dist
        }
    }
}