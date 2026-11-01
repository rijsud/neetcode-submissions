func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	} )

	prevEnd := intervals[0][1]
	res := 0
	for _, interval := range intervals[1:] {
		if interval[0] < prevEnd {
			res++
			prevEnd = min(interval[1], prevEnd)
		} else {
			prevEnd = interval[1]
		}
	}

	return res
}
