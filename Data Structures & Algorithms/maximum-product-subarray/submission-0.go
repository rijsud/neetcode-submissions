func maxProduct(nums []int) int {
    res, curMin, curMax := nums[0], 1, 1
	for _, num := range nums {
		curMin, curMax = min(num * curMax, num * curMin, num), max(num * curMax, num * curMin, num)
		res = max(curMax, res)
	}
	return res
}
