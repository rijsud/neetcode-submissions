func maxSlidingWindow(nums []int, k int) []int {
    l := 0 // curMax always heapQueue[0]
	q := []int{}
	output := []int{}
	// for i := 0; i < k; i++ {
	// 	heapQueue = append(heapQueue, i)
	// 	if nums[heapQueue[0]] > nums[i] {
	// 		heapQueue = append(heapQueue, i)
	// 	} else {
	// 		heapQueue[len(heapQueue)-1] = i
	// 	}
	// }
	for r := 0; r < len(nums); r++ {
		for len(q) > 0 && nums[q[len(q)-1]] < nums[r] {
			q = q[:len(q)-1]
		}
		q = append(q, r)
		if l > q[0] {
			q = q[1:]
		}
		if (r+1) >= k {
			output = append(output,nums[q[0]])
			l++
		}
	}

	return output
}
