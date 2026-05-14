func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numSet := make(map[int]int)
	for index, number := range nums1 {
		numSet[number] = index
	}

	res := make([]int, len(nums1))
	for i := range res {
		res[i] = -1
	}

	stack := []int{}
	for _, curNum := range nums2 {
		for len(stack) > 0 && curNum > stack[len(stack)-1] {
			val := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			index := numSet[val]
			res[index] = curNum
		}
		if _, ok := numSet[curNum]; ok {
			stack = append(stack, curNum)
		}
	}

	return res
}
