func rob(nums []int) int {
    rob1, rob2 := 0, 0
	for i := 0; i < len(nums); i++ {
		rob1, rob2 = rob2, max(rob1 + nums[i], rob2)
	}
	return rob2
}
