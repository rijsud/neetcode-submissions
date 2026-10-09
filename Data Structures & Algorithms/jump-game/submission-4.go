// func canJump(nums []int) bool {
// 	i := 0
//     for i < len(nums) && nums[i] != 0 {
// 		maxJump = nums[i]
// 	}
// 	if i >= len(nums) - 1 {return true}
// 	return false
// }
func canJump(nums []int) bool {
    goal := len(nums) - 1
    for i := len(nums) - 2; i > -1; i-- {
        if i + nums[i] >= goal {
            goal = i
        }
    }
    
    return goal == 0
}