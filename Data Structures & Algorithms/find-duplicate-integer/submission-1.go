func findDuplicate(nums []int) int {
    fast, slow := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {break}
	}

	slow = 0
	for {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {break}
	}
	
	return slow
}
