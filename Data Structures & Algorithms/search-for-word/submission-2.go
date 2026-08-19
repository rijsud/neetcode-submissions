func exist(board [][]byte, word string) bool {
    rows, cols := len(board), len(board[0])
    
    var backtrack func(int, int, int) bool
    backtrack = func(r int, c int, i int) bool {
        if i == len(word) {return true}
        if r == rows || c == cols || r < 0 || c < 0 || 
        board[r][c] != word[i] || board[r][c] == '#' {return false}
        temp := board[r][c]
        board[r][c] = '#'
        res :=  backtrack(r + 1, c    , i + 1) ||
                backtrack(r - 1, c    , i + 1) ||
                backtrack(r    , c + 1, i + 1) ||
                backtrack(r    , c - 1, i + 1)
        board[r][c] = temp
        return res
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            if backtrack(r, c, 0) {return true}
        }
    }
    return false
}
