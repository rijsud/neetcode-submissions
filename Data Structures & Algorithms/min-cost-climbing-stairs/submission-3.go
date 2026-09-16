func minCostClimbingStairs(cost []int) int {
	n := len(cost)
    cache := make([]int, n)
	for i  := range cache {cache[i] = -1}

	var dfs func(step int) int
	dfs = func(step int) int {
		if step >= n {return 0}
		if cache[step] != -1 {return cache[step]}
		cache[step] = cost[step] + min(dfs(step + 1), dfs(step + 2))
		return cache[step]
	}

	return min(dfs(0), dfs(1))
}
