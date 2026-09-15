func rob(nums []int) int {
	one := nums[0] // what if the length of the array was 0?
    if len(nums) == 1 {return one}
	two := max(one, nums[1]) // two represents the current maximum, so needs to be max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		one, two = two, max(one + nums[i], two)
	}
	return two
}
