func maxSlidingWindow(nums []int, k int) []int {
    output := make([]int, 0, len(nums)-k+1)

    for i := 0; i <= len(nums)-k; i++ {
        maxi := nums[i]
        for j := i; j < i+k; j++ {
            if nums[j] > maxi {
                maxi = nums[j]
            }
        }
        output = append(output, maxi)
    }

    return output
}
