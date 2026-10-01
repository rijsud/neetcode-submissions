func findTargetSumWays(nums []int, target int) int {
	dp := map[[2]int]int{}
    var dfs func(i, sum int) int
	dfs = func(i, sum int) int {
		if i == len(nums) {
			if sum == target { return 1 }
			return 0
		}
		key := [2]int{i, sum}
		if val, found := dp[key]; found { return val }

		dp[key] = dfs(i + 1, sum - nums[i]) + dfs(i + 1, sum + nums[i])
		return dp[key]
	}

	return dfs(0, 0)
}
