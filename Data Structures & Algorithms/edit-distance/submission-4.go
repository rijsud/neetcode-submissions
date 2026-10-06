func minDistance(word1, word2 string) int {
    m, n := len(word1), len(word2)
    if m < n {
        word1, word2 = word2, word1
        m, n = n, m
    }

    dp := make([]int, n+1)

    for j := 0; j <= n; j++ {
        dp[j] = n - j
    }

    for i := m - 1; i >= 0; i-- {
    	nextDp := make([]int, n+1)
        nextDp[n] = m - i
        for j := n - 1; j >= 0; j-- {
            if word1[i] == word2[j] {
                nextDp[j] = dp[j+1]
            } else {
                nextDp[j] = 1 + min(dp[j], nextDp[j+1], dp[j+1])
            }
        }
        dp = nextDp
    }

    return dp[0]
}