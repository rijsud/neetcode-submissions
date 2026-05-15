func maxProfit(prices []int) int {
	minNum, maxNum := prices[0], prices[0]
	profit := 0

	for _, num := range prices {
		if num < minNum {
			minNum = num
			maxNum = num
		}
		if num > maxNum {maxNum = num}
		profit = max(profit, maxNum - minNum)
	}

	return profit
}
