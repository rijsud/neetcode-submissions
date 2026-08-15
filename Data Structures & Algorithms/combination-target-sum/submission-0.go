func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    n := len(nums)
    subset := []int{}
    var dfs func(int, int)
    dfs = func(i int, sum int) {
        if sum > target || i >= n { return }
        if target == sum {
            subcopy := make([]int, len(subset))
            copy(subcopy, subset)
            res = append(res, subcopy)
            return
        }
        subset = append(subset, nums[i])
        dfs(i, sum + nums[i])
        subset = subset[:len(subset)-1]
        dfs(i + 1, sum)
    }
    dfs(0, 0)
    return res
}
