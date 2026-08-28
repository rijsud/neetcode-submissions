func islandsAndTreasure(grid [][]int) {
    rows, cols := len(grid), len(grid[0])
    directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

    var bfs func(int, int)
    bfs = func(r, c int) {
        q := [][3]int{{r, c, 0}}
        for len(q) > 0 {
            row, col, dist := q[0][0], q[0][1], q[0][2]
            q = q[1:]

            for _, dir := range directions {
                newRow, newCol := row + dir[0], col + dir[1]
                if newRow < 0 || newCol < 0 ||
                newRow == rows || newCol == cols ||
                grid[newRow][newCol] == -1 {continue}
                if grid[newRow][newCol] > dist + 1 {
                    q = append(q, [3]int{newRow, newCol, dist + 1})
                    grid[newRow][newCol] = dist + 1
                }
            }
        }
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if grid[r][c] == 0 {bfs(r, c)}
        }
    }
}
