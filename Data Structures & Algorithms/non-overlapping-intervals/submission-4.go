func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	} )

	prevEnd := intervals[0][1]
	res := 0
	for _, interval := range intervals[1:] {
		if prevEnd > interval[0] {
			res++
		} else {
			prevEnd = interval[1]
		}
	}

	return res	
}
