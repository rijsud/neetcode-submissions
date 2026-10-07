func maxCoins(nums []int) int {
    nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {dp[i] = make([]int, n)}

	var dfs func(int, int) int
	dfs = func(l, r int) int {
		if l > r {return 0}
		if dp[l][r] > 0 {return dp[l][r]}

		for i := l; i <= r; i++ {
			coins := nums[i] * nums[l-1] * nums[r+1]
			coins += dfs(l, i-1) + dfs(i+1, r)
			dp[l][r] = max(dp[l][r], coins)
		}

		return dp[l][r]
	}

	return dfs(1, n-2)
}
