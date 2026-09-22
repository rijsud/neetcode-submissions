func lengthOfLIS(nums []int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums) + 1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(i, j int) int
    dfs = func(i, j int) int {
        if i == len(nums) {
            return 0
        }
		if dp[i][j + 1] != -1 { return dp[i][j + 1] }

        LIS := dfs(i + 1, j) // not include

        if j == -1 || nums[j] < nums[i] {
            LIS = max(LIS, 1 + dfs(i + 1, i)) // include
        }

		dp[i][j + 1] = LIS
        return LIS
    }

    return dfs(0, -1)
}
