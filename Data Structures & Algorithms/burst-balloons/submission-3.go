func maxCoins(nums []int) int {
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {dp[i] = make([]int, n)}

	for l := n - 2; l >= 1; l-- {
		for r := 1; r <= n-2; r++ {
			for i := l; i <= r; i++ {
				coins := nums[l-1] * nums[i] * nums[r+1]
				coins += dp[i+1][r] + dp[l][i-1]
				dp[l][r] = max(dp[l][r], coins)
			}
		}
	}

	return dp[1][n-2]
}
