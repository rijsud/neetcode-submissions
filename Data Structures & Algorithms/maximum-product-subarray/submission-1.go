func maxProduct(nums []int) int {
	res, curMin, curMax := nums[0], 1, 1
	for _, num := range nums {
		temp := curMin * num
		curMin = min(num, curMin * num, curMax * num)
		curMax = max(num, temp, curMax * num)
		res = max(res, curMax)
	}
	return res
}
