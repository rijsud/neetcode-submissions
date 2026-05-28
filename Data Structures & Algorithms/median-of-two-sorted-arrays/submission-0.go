func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sliceA, sliceB := nums1, nums2
	totalLength := len(nums1) + len(nums2)
	half :=  (totalLength + 1) / 2

	if len(sliceB) < len(sliceA) {
		sliceA, sliceB = sliceB, sliceA
	}

	l, r := 0, len(sliceA)
	for {
		midA := l + ((r - l) / 2)
		midB := half - midA

		leftA, rightA := math.MinInt64, math.MaxInt64
		leftB, rightB := math.MinInt64, math.MaxInt64

		if midA > 0 {leftA = sliceA[midA-1]}
		if (midA + 1) <= len(sliceA) {rightA = sliceA[midA]}
		if midB > 0 {leftB = sliceB[midB-1]}
		if (midB + 1) <= len(sliceB) {rightB = sliceB[midB]}

		if leftA <= rightB && leftB <= rightA {
			if totalLength % 2 != 0 {
				return float64(max(leftA, leftB))
			} else {
				return float64(max(leftA, leftB)+min(rightA, rightB)) / 2
			}
		} else if leftA > rightB {
			r = midA - 1
		} else {
			l = midA + 1
		}
	}

	return -1
}
