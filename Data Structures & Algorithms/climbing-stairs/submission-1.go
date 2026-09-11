// func climbStairs(n int) int {
//     cache := make([]int, n)
// 	for i := range cache {cache[i] = -1}

// 	var dfs func(int) int
//     dfs = func(steps int) int {
//         if steps >= n {
//             if steps == n {
//                 return 1
//             }
//             return 0
//         }
// 		if cache[steps] != -1 {return cache[steps]}

//         cache[steps] = dfs(steps + 1) + dfs(steps + 2)
// 		return cache[steps]
//     }

//     return dfs(0)
// }

func climbStairs(n int) int {
	cache := make([]int, n + 1)
	for i := range cache {cache[i] = -1}

    var dfs func(int) int
	dfs = func(steps int) int {
		if steps == 1 {return 1}
		if steps == 2 {return 2}
		
		if cache[steps] != -1 {return cache[steps]}
		cache[steps] = dfs(steps - 1) + dfs(steps - 2)
		return cache[steps]
	}

	return dfs(n)
}