func validTree(n int, edges [][]int) bool {
    allEdges := make([][]int, n)
	for _, edge := range edges {
		allEdges[edge[0]] = append(allEdges[edge[0]], edge[1])
		allEdges[edge[1]] = append(allEdges[edge[1]], edge[0])
	}

    visit := map[int]bool{}

	var dfs func(int, int) bool
	dfs = func(parent int, node int) bool {
		if visit[node] {return false}
		visit[node] = true
		for _, nei := range allEdges[node] {
			if nei == parent {continue}
			if !dfs(node, nei) {return false}
		}
		return true
	}

	return dfs(-1, 0) && len(visit) == n
}
