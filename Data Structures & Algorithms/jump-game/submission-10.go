func canJump(nums []int) bool {
    goal := len(nums)-1
	for i := len(nums)-2; i > -1; i-- {
		if goal <= nums[i] + i {
			goal = i
		}
	}
	return goal == 0
}
