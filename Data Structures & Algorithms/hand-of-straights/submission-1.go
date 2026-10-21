func isNStraightHand(hand []int, groupSize int) bool {
    sort.Ints(hand)
	found := map[int]int{}
	for _, num := range hand {
		found[num]++
	}

	for _, num := range hand {
		if found[num] > 0 {
			for i := num; i < num+groupSize; i++ {
				found[i]--
				if found[i] < 0 {
					return false
				}
			}
		}
	}
	return true
}
