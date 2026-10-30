func insert(intervals [][]int, newInterval []int) [][]int {
    res := [][]int{}
	for i := range intervals {
		if newInterval[1] < intervals[i][0]  {
			res = append(res, newInterval)
			return append(res, intervals[i:]...)
		} else if newInterval[0] > intervals[i][1] {
			res = append(res, intervals[i])
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}
	return append(res, newInterval)
}
