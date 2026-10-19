func canCompleteCircuit(gas []int, cost []int) int {
	sum, res, total := 0, 0, 0
    for i := range gas {
		diff := gas[i] - cost[i]
		total += diff
		sum += diff
		if total < 0 {
			total = 0
			res = i + 1
		}
	}
	if sum < 0 {return -1} else {return res}
}
