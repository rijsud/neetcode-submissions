func isValidSudoku(board [][]byte) bool {
    // dupliMap := make(map[byte][][2]int)
    // for rowIndex, row := range board {
    //     for columnIndex, number := range row {
    //         if number == '.' { continue }
    //         dupliMap[number] = append(dupliMap[number], [2]int{rowIndex, columnIndex})
    //     }
    // }

    // for key, value := range dupliMap {
    //     for i, coordinate := range value {
    //         if 
    //     }
    // } 
    rows := make([]map[byte]bool, 9)
    cols := make([]map[byte]bool, 9)
    squares := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        rows[i] = make(map[byte]bool)
        cols[i] = make(map[byte]bool)
        squares[i] = make(map[byte]bool)
    }

    for r := 0; r < 9; r++ {
        for c := 0; c < 9; c++ {
            if board[r][c] == '.' {
                continue
            }
            val := board[r][c]
            squareIdx := (r/3)*3 + c/3

            if rows[r][val] || cols[c][val] ||
               squares[squareIdx][val] {
                return false
            }

            rows[r][val] = true
            cols[c][val] = true
            squares[squareIdx][val] = true
        }
    }

    return true    
}
