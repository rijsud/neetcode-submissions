func minEatingSpeed(piles []int, h int) int {
	l, r := 1, 0
	for _, pile := range piles {
		if pile > r {
			r = pile
		}
	}

	largest := r
	if h == len(piles) {
		return largest
	}

	for l <= r {
		mid := l + ((r-l)/2)
		totalH := 0

		for _, bananas := range piles {
			totalH += int(math.Ceil(float64(bananas)/float64(mid)))
		}

		if totalH <= h {
			r = mid - 1
			largest = mid
		} else {
			l = mid + 1
		}
	}

	return largest
}