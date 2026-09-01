func canFinish(numCourses int, prerequisites [][]int) bool {
    preMap := make([][]int, numCourses)
	for i := range preMap {
		preMap[i] = []int{}
	}
	for _, val := range prerequisites {
		preMap[val[0]] = append(preMap[val[0]], val[1])
	}

	visit := map[int]bool{}
	var dfs func(int) bool
	dfs = func(crs int) bool {
		if visit[crs] {return false}
		if len(preMap[crs]) == 0 {return true}

		visit[crs] = true
		for _, pre := range preMap[crs] {
			if !dfs(pre) {return false}
		}
		visit[crs] = false
		preMap[crs] = []int{}
		return true
	}

	for crs := 0; crs < numCourses; crs++ {
		if !dfs(crs) {return false}
	}

	return true
}
