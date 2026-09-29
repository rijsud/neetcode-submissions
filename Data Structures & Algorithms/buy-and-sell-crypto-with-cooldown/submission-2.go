type mapKey struct {
	index int
	ifBuy bool
}

func maxProfit(prices []int) int {
	cache := map[mapKey]int{}
    var dfs func(int, bool) int
	dfs = func(i int, buying bool) int {
		if i >= len(prices) {return 0}
		key := mapKey{index : i, ifBuy : buying}
		if val, found := cache[key]; found {
			return val
		}

		cooldown := dfs(i + 1, buying)
		if buying {
			buy := dfs(i + 1, false) - prices[i]
			cache[key] = max(buy, cooldown)
		} else {
			sell := dfs(i + 2, true) + prices[i]
			cache[key] = max(sell, cooldown)
		}

		return cache[key]
	}

	return dfs(0, true)
}
