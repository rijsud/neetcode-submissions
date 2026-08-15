func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    sort.Ints(nums)
    n := len(nums)
    subset := []int{}

    var dfs func(int, int)
    dfs = func(i int, sum int) {
        if target == sum {
            subcopy := make([]int, len(subset))
            copy(subcopy, subset)
            res = append(res, subcopy)
            return
        }
        for j := i; j < n; j++ {
            if sum + nums[j] > target { return }
            subset = append(subset, nums[j])
            dfs(j, sum + nums[j])
            subset = subset[:len(subset)-1]
        }
    }
    
    dfs(0, 0)
    return res
}
