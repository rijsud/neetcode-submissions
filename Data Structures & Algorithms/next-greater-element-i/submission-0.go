func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := []int{}
	for _, num1 := range nums1 {
		appendNumber := -1
		index := 0
		for j, num2 := range nums2 {
			if num2 == num1 {
				index = j
				break
			}
		}
		for j := index; j < len(nums2); j++ {
			if nums2[j] > num1 {
				appendNumber = nums2[j]
				break
			}
		}
		res = append(res, appendNumber)
	}
	return res
}
