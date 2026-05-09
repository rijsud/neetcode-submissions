func longestConsecutive(nums []int) int {
	longest := 0
	numSet := make(map[int]struct{})

	for _, number := range nums {
		numSet[number] = struct{}{}
	}

	for number := range numSet {
		if _, found := numSet[number - 1] ; !found {
			length := 1
			// for _, ok := numSet[number + length]; ok; _, ok = numSet[number + length] {
			// 	length++
			// }
			for {
				if _, exists := numSet[number+length]; exists{
					length++
				} else {
					break
				}
			}
			longest = max(longest, length)
		}
	}

	return longest
}
