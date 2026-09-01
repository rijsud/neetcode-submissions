func findOrder(numCourses int, prerequisites [][]int) []int {
    preMap := make([][]int, numCourses)
	cycle, seen := map[int]bool{}, map[int]bool{}
	res := []int{}

	for _, pre := range prerequisites {
		preMap[pre[0]] = append(preMap[pre[0]], pre[1])
	}

	var dfs func(int) bool
	dfs = func(crs int) bool {
		if cycle[crs] {return false}
		if seen[crs] {return true}

		cycle[crs] = true
		for _, pre := range preMap[crs] {
			if !dfs(pre) {return false}
		}
		cycle[crs] = false
		seen[crs] = true
		res = append(res, crs)
		return true
	}

	for c := 0; c < numCourses; c++ {
		if !dfs(c) {return []int{}}
	}

	return res
}
