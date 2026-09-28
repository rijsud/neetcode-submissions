func uniquePaths(m int, n int) int {
    dp := make([]int, n)
    for i := range dp {dp[i] = 1}

    for i := m - 2; i > -1; i-- {
        for j := n - 2; j > -1; j-- {
            dp[j] += dp[j+1]
        }
    }

    return dp[0]
}
