func canPartition(nums []int) bool {
	totalSum := 0
	for i := range nums {totalSum += nums[i]}
	if totalSum % 2 != 0 {return false}
	totalSum /= 2

	dp := map[int]bool{ 0 : true }
	for _, num := range nums {
		nextDP := map[int]bool{}
		for n := range dp {
			if num + n == totalSum {return true}
			nextDP[n] = true
			if num + n > totalSum {continue}
			nextDP[n+num] = true
		}
		dp = nextDP
	}
	return false
}
