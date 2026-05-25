func searchMatrix(matrix [][]int, target int) bool {
	rows, columns := len(matrix), len(matrix[0])
	l, r := 0, (rows * columns) - 1

	for l <= r {
		mid := l + ((r-l)/2)
		row, col := mid/columns, mid%columns
		if matrix[row][col] > target {
			r = mid - 1
		} else if matrix[row][col] < target {
			l = mid + 1
		} else {
			return true
		}
	}
	return false
}
