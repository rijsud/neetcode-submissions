func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
    sort.Ints(nums)
	n := len(nums)
	subset := []int{}

	var backtrack func(int)
	backtrack = func(i int) {
		if i == n {
			subcopy := make([]int, len(subset))
			copy(subcopy, subset)
			res = append(res, subcopy)
			return
		}
        subset = append(subset, nums[i])
        backtrack(i + 1)
        subset = subset[:len(subset)-1]

        for i + 1 < n && nums[i] == nums[i+1] {i++}
        backtrack(i + 1)
	}

	backtrack(0)

	return res
}