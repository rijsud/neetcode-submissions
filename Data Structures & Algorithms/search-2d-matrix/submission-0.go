func searchMatrix(matrix [][]int, target int) bool {
	top, bottom := 0, len(matrix)-1
	for top <= bottom {
		midRow := top + ((bottom-top)/2)
		l, r := 0, len(matrix[midRow])-1
		if matrix[midRow][r] < target {
			top = midRow + 1
		} else if matrix[midRow][l] > target {
			bottom = midRow - 1
		} else {
			for l <= r {
				mid := l + ((r-l)/2)
				if matrix[midRow][mid] < target {
					l = mid + 1
				} else if matrix[midRow][mid] > target {
					r = mid - 1
				} else {
					return true
				}
			}
			return false
		}
	}

	return false
}
