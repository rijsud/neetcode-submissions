func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
    directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    visit := map[[2]int]bool{}
    q := [][2]int{}


    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 0 {
                q = append(q, [2]int{r,c})
                visit[[2]int{r,c}] = true
            }
        }
    }

    dist := 0
    for len(q) > 0 {
        qLen := len(q)
        for i := 0; i < qLen; i++ {
            r, c := q[0][0], q[0][1]
            q = q[1:]
            grid[r][c] = dist

            for _, dir := range directions {
                nr, nc := r + dir[0], c + dir[1]
                if nr < 0 || nc < 0 || nr == rows || nc == cols ||
                visit[[2]int{nr,nc}] || grid[nr][nc] == -1 {continue}
                visit[[2]int{nr,nc}] = true
                q = append(q, [2]int{nr,nc})
            }
        }
        dist++
    }
}
