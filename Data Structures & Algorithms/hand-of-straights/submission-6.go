func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize != 0 {return false}

	freq := map[int]int{}
	for _, num := range hand {
		freq[num]++
	}

	for _, num := range hand {
		start := num
		for freq[start-1] > 0 {
			start--
		}

		for start <= num {
			for freq[start] > 0 {
				for i := start; i < start+groupSize; i++ {
					freq[i]--
					if freq[i] < 0 {
						return false
					}
				}
			}
			start++
		}
		
	}
	return true
}
