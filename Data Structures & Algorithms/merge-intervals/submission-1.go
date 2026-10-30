func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
	res := [][]int{intervals[0]}
    for _, interval := range intervals[1:] {
		if res[len(res)-1][1] >= interval[0] {
			res[len(res)-1] = []int{
				min(res[len(res)-1][0], interval[0]),
				max(res[len(res)-1][1], interval[1]),
			}
		} else {
			res = append(res, interval)
		}
	}

	return res
}
