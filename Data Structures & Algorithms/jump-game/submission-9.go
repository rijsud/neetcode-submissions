func canJump(nums []int) bool {
    reach := 0
	for i, num := range nums[:len(nums)-1] {
		if reach < i {break}
		reach = max(reach, num + i)
	}
	return reach >= len(nums)-1
}
