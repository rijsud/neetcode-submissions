func maxCoins(nums []int) int {
    nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var dfs func(int, int) int
	dfs = func(l, r int) int {
		if l > r {return 0}
		if dp[l][r] != -1 {return dp[l][r]}
		maxCoins := 0
		for i := l; i <= r; i++ {
			coins := nums[l-1] * nums[i] * nums[r+1]
			coins += dfs(i+1, r) + dfs(l, i-1)
			maxCoins = max(maxCoins, coins)
		}
		dp[l][r] = maxCoins
		return maxCoins
	}

	return dfs(1, n-2)
}
