func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
    path := map[[2]int]bool{}

    var backtrack func(int, int, int) bool
    backtrack = func(r, c, i int) bool {
        if i == len(word) {return true}
        if r < 0 || c < 0 || r >= rows || c >= cols ||
        word[i] != board[r][c] || path[[2]int{r,c}] {return false}

        path[[2]int{r,c}] = true
        res := backtrack(r + 1, c, i + 1) || backtrack(r, c + 1, i + 1) ||
        backtrack(r - 1, c, i + 1) || backtrack(r, c - 1, i + 1)
        delete(path, [2]int{r,c})
        return res
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if backtrack(r, c, 0) {return true}
        }
    }

    return false
}
