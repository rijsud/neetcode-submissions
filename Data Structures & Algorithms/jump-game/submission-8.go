func canJump(nums []int) bool {
	n := len(nums)
	goal := n - 1
    for i := n - 2; i > -1; i-- {
		if nums[i] + i >= goal {
			goal = i
		}
	}
	return goal == 0
}
