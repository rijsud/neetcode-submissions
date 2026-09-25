func lengthOfLIS(nums []int) int {
	n := len(nums)
    dp := make([]int, n)
	for i := range dp {dp[i] = 1}
	maxLen := 1

	for i := n - 2; i > -1; i-- {
		for j := i; j < n; j++ {
			if nums[j] > nums[i] {
				dp[i] = max(dp[i], 1 + dp[j])
			}
		}
		maxLen = max(dp[i], maxLen)
	}

	return maxLen
}
