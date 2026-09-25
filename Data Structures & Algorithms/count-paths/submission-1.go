func uniquePaths(m int, n int) int {
    if m == 1 || n == 1 {return 1}
    if m < n {
        m, n = n, m
    }
    
    res, j := 1, 1
    for i := m; i < m + n - 1; i++ {
        res *= i
        res /= j
        j++
    }

    return res
}
