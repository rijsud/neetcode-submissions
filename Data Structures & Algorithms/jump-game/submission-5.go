func canJump(nums []int) bool {
    goal := len(nums)-1
    for i := len(nums)-2; i > -1; i-- {
        if nums[i] + i >= goal {
            goal = i
        }
    }

    return goal == 0
}
