func maxProduct(nums []int) int {
    res, curMin, curMax := nums[0], 1, 1
	for _, num := range nums {
		curMin, curMax = min(num, curMin*num, curMax*num), max(num, curMin*num, curMax*num)
		res = max(res, curMax)
	}
	return res
}
