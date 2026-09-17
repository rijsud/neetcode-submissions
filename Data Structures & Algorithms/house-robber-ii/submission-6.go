func rob(nums []int) int {
    return max(nums[0], max(helper(nums[1:]),
                            helper(nums[:len(nums)-1])))
}

func helper(nums []int) int {
    rob1, rob2 := 0, 0
    for _, num := range nums {
		rob1, rob2 = rob2, max(rob1 + num, rob2)
    }
    return rob2
}