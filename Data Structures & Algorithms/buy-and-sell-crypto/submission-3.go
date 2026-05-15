func maxProfit(prices []int) int {
	minIndex := 0
	maxPft := 0

	for i := 1; i < len(prices); i++ {
		if prices[minIndex] < prices[i] {
			profit := prices[i] - prices[minIndex]
			maxPft = max(maxPft, profit)
		} else {
			minIndex = i
		}
	}
	return maxPft
}
