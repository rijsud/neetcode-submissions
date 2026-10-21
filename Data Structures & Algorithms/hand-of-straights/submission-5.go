func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand) % groupSize != 0 {return false}
	
	count := map[int]int{}
	for _, num := range hand {
		count[num]++
	}

	for _, num := range hand {
		start := num
		for count[start-1] > 0 {
			start--
		}
		for start <= num {
			for count[start] > 0 {
				for i := start; i < start+groupSize; i++ {
					if count[i] == 0 {
						return false
					}
					count[i]--
				}
			}
			start++
		}
	}
	return true
}
