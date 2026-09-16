func rob(nums []int) int {
	n := len(nums)
	if n == 1 {return nums[0]}
    rob1, rob2 := 0, 0
	for i := 0; i < n - 1; i++ {
		rob1, rob2 = rob2, max(rob1 + nums[i], rob2)
	}
	maxRob := rob2
	rob1, rob2 = 0, 0
	for i := n - 1; i > 0; i-- {
		rob1, rob2 = rob2, max(rob1 + nums[i], rob2)
	}
	return max(maxRob, rob2)
}
