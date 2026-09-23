func canPartition(nums []int) bool {
	totalSum := 0
	for i := range nums {totalSum += nums[i]}
	if totalSum % 2 != 0 {return false}
	totalSum /= 2

	dp := map[int]bool{0 : true}

	for i := len(nums) - 1; i > -1; i-- {
		nextDP := map[int]bool{}
		for t := range dp {
			if nums[i] + t == totalSum {
				return true
			}
			nextDP[nums[i] + t] = true
			nextDP[t] = true
		}
		dp = nextDP
	}

	return false
}
