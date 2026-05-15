func maxProfit(prices []int) int {
	minNum:= prices[0]
	profit := 0

	for _, num := range prices {
		if num > minNum {
			profit = max(profit, num - minNum)
		} else {
			minNum = num
		}
	}

	return profit
}
