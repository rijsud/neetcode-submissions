func findTargetSumWays(nums []int, target int) int {
    dp := map[int]int{}
	dp[0] = 1
	for _, num := range nums {
		nextDP := map[int]int{}
		for sum, count := range dp {
			nextDP[sum - num] += count
			nextDP[sum + num] += count
		}
		dp = nextDP
	}
	return dp[target]
}
