func coinChange(coins []int, amount int) int {
	cache := map[int]int{0 : 0}

    var dfs func(int) int
	dfs = func(amount int) int {
		if amount == 0 {return 0}
		if val, ok := cache[amount]; ok {return val}
		res := math.MaxInt32
		for j := 0; j < len(coins); j++ {
			if amount - coins[j] >= 0 {
				res = min(res, 1 + dfs(amount - coins[j]))
			}
		}
		cache[amount] = res
		return cache[amount]
	}

	minCoins := dfs(amount)
	if minCoins >= math.MaxInt32 {
        return -1
    }
    return minCoins
}
