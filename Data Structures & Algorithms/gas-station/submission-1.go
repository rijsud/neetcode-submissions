func canCompleteCircuit(gas []int, cost []int) int {
    sum := 0
	for i := range gas {
		sum += gas[i] - cost[i]
	}
	if sum < 0 {return -1}

	total, res := 0, 0
	for i := range gas {
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			res = i + 1
		}
	}

	return res
}
