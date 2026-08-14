func subsets(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	subset := []int{}

	var backtrack func(int)
	backtrack = func(i int) {
		if i >= n {
			subcopy := make([]int, len(subset))
			copy(subcopy, subset)
			res = append(res, subcopy)
			return
		}
        backtrack(i + 1)
        subset = append(subset, nums[i])
        backtrack(i + 1)
        subset = subset[:len(subset)-1]
	}

	backtrack(0)

	return res
}
