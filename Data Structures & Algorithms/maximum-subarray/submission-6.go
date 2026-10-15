func maxSubArray(nums []int) int {
	maxSum, curSum := nums[0], 0
	for _, num := range nums {
		if curSum < 0 {curSum = 0}
		curSum += num
		if maxSum < curSum {maxSum = curSum}
	}
	return maxSum
}
