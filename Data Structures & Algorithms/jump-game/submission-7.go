func canJump(nums []int) bool {
    reach := 0
	for i, num := range nums {
		if reach < i {break}
		if num + i > reach {reach = num + i}
	}
	return reach >= len(nums)-1
}
