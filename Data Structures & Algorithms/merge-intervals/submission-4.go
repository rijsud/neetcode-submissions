func merge(intervals [][]int) [][]int {
    mp := map[int]int{}
	for _, interval := range intervals {
		mp[interval[0]]++
		mp[interval[1]]--
	}

	res := [][]int{}
	interval := []int{}
	keys := []int{}
	check := 0

	for key := range mp {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		if len(interval) == 0 {
			interval = append(interval, key)
		}
		check += mp[key]
		if check == 0 {
			interval = append(interval, key)
			res = append(res, interval)
			interval = []int{}
		}
	}

	return res
}
