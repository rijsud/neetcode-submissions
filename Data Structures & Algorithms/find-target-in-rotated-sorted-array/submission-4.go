func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l < r {
		mid := l + ((r - l) / 2)
		if nums[mid] < nums[r]{
			r = mid
		} else {
			l = mid + 1
		}
	}

	pivot := l
	l, r = 0, len(nums)-1

	if target >= nums[pivot] && target <= nums[r] {
        l = pivot
    } else {
        r = pivot - 1
    }

	for l <= r {
		mid := l + ((r - l) / 2)
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
