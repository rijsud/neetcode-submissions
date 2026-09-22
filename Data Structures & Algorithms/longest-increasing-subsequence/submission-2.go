func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
	for i := range dp {dp[i] = 1}
	maxLen := dp[0]

	for i := len(nums)-2; i > -1; i-- {
		for j := i; j < len(nums); j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1 + dp[j])
			}
		}
		maxLen = max(dp[i], maxLen)
	}

	return maxLen
}
