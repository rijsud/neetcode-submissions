func maxProfit(prices []int) int {
	minNum:= prices[0]
	profit := 0

	for _, num := range prices {
		if num < minNum {
			minNum = num
		}
		profit = max(profit, num - minNum)
	}

	return profit
}
