func lengthOfLIS(nums []int) int {
    dp := []int{}
	dp = append(dp, nums[0])

	for i := 1; i < len(nums); i++ {
		if dp[len(dp)-1] < nums[i] {
			dp = append(dp, nums[i])
			continue
		}

		idx := lowerBound(dp, nums[i])
		dp[idx] = nums[i]
	}

	return len(dp)
}

func lowerBound(dp []int, target int) int {
	l, r := 0, len(dp)-1
	ans := -1
	for l <= r {
		mid := l + ((r-l)/2)
		if dp[mid] >= target {
			r = mid - 1
			ans = mid
		} else {
			l = mid + 1
		}
	}
	return ans
}
