func pacificAtlantic(heights [][]int) [][]int {
    rows, cols := len(heights), len(heights[0])
	pacific := map[[2]int]bool{}
	atlantic := map[[2]int]bool{}
	res := [][]int{}
	directions := [][2]int{{0,1},{0,-1},{1,0},{-1,0}}

	var dfs func(int, int, map[[2]int]bool, int)
	dfs = func(r, c int, visit map[[2]int]bool, prevHeight int) {
		if r < 0 || c < 0 || r == rows || c == cols || 
		visit[[2]int{r,c}] || heights[r][c] < prevHeight {return}

		visit[[2]int{r,c}] = true
		for _, dir := range directions {
			dfs(r + dir[0], c + dir[1], visit, heights[r][c])
		}
	}

	for c := 0; c < cols; c++ {
		dfs(0, c, pacific, heights[0][c])
		dfs(rows-1, c, atlantic, heights[rows-1][c])
	}

	for r := 0; r < rows; r++ {
		dfs(r, 0, pacific, heights[r][0])
		dfs(r, cols-1, atlantic, heights[r][cols-1])
	}

	for key, _ := range pacific {
		if atlantic[key] {res = append(res, []int{key[0], key[1]})}
	}

	return res
}
