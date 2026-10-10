func canJump(nums []int) bool {
    jumpLen := nums[0]
    for i := 1; i < len(nums); i++ {
        if jumpLen < 1 {return false}
        jumpLen--
        if nums[i] > jumpLen {jumpLen = nums[i]}
    }
    return true
}
