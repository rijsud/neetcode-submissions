func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], nums[0]
	for _, num := range nums[1:] {
		curSum = max(num, curSum+num)
		maxSum = max(curSum, maxSum)
	}
	return maxSum
}
