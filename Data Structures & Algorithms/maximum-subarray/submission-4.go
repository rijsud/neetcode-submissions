func maxSubArray(nums []int) int {
    curSum, maxSum := 0, nums[0]
	for _, num := range nums {
		curSum = max(num, curSum + num)
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}
