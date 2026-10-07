func maxCoins(nums []int) int {
    nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	n := len(nums)

	dp := make([][]int, n)
	for i := range dp {dp[i] = make([]int, n)}

	for l := n; l > 0; l-- {
		for r := 0; r < n-1; r++ {
			for i := l; i <= r; i++ {
				coins := nums[i] * nums[l-1] * nums[r+1]
				coins += dp[l][i-1] + dp[i+1][r]
				dp[l][r] = max(dp[l][r], coins)
			}
		}
	}

	return dp[1][n-2]
}
