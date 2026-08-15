func combinationSum2(candidates []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(candidates)
    n := len(candidates)
    subset := []int{}
    
    var dfs func(int, int)
    dfs = func(i int, sum int) {
        if sum == target {
            subcopy := make([]int, len(subset))
            copy(subcopy, subset)
            res = append(res, subcopy)
            return
        }
        if i >= n || sum > target {return}
        subset = append(subset, candidates[i])
        dfs(i + 1, sum + candidates[i])
        subset = subset[:len(subset)-1]
        for i + 1 < n && candidates[i + 1] == candidates[i] {
            i++
        } 
        dfs(i + 1, sum)
    }

    dfs(0, 0)
    return res
}
