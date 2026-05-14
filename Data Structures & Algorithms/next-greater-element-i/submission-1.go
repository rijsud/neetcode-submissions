func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numSet := make(map[int]int)
	res := []int{}
	for index, number := range nums2 {
		numSet[number] = index
	}

	for _, number := range nums1 {
		index, _ := numSet[number]
		appendNumber := -1
		for i := index; i < len(nums2); i++ {
			if nums2[i] > number {
				appendNumber = nums2[i]
				break
			} 
		}
		res = append(res, appendNumber)
	}

	return res
}
