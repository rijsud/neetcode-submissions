func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    INF := math.MaxInt32
	prices := make([]int, n)
	for i := range prices { prices[i] = INF	}
	prices[src] = 0
	
	for i := 0; i <= k; i++ {
		temp := make([]int, n)
		copy(temp, prices)
		for _, flight := range flights {
			s, d, p := flight[0], flight[1], flight[2]
			if prices[s] == INF {continue}
			if prices[s] + p < temp[d] {temp[d] = prices[s] + p}
		}
		prices = temp
	}
	
	if prices[dst] == INF {return -1}
	return prices[dst]
}
