func isNStraightHand(hand []int, groupSize int) bool {
    if len(hand)%groupSize != 0 {return false}

	sort.Ints(hand)
	freq := map[int]int{}
	for _, num := range hand {
		freq[num]++
	}

	for _, num := range hand {
		if freq[num] > 0 {
			for i := num; i < num+groupSize; i++ {
				freq[i]--
				if freq[i] < 0 {
					return false
				}
			}
		}
	}
	return true
}
