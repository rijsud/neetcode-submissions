func findItinerary(tickets [][]string) []string {
    adj := map[string][]string{}
	for _, ticket := range tickets {
		src, dst := ticket[0], ticket[1]
		adj[src] = append(adj[src], dst)
	}

	for src := range adj {
		sort.Strings(adj[src])
	}

	res := []string{}
	var dfs func(string)
	dfs = func(src string) {
		for len(adj[src]) > 0 {
			to := adj[src][0]
			adj[src] = adj[src][1:]
			dfs(to)
		}
		res = append(res, src)
	}

	dfs("JFK")
	for i, j := 0, len(res) - 1; i < j; i, j = i + 1, j - 1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}
