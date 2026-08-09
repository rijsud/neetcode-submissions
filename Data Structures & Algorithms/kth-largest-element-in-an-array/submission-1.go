// Quick Select algorithm

func findKthLargest(nums []int, k int) int {
    k = len(nums) - k
    
    var quickSelect func(int, int) int
    quickSelect = func(left, right int) int {
        p := left
        for i := left; i < right; i++ {
            if nums[i] <= nums[right] {
                nums[i], nums[p] = nums[p], nums[i]
                p++
            }
        }
        nums[p], nums[right] = nums[right], nums[p]

        if k < p {
            return quickSelect(left, p - 1)
        } else if k > p {
            return quickSelect(p + 1, right)
        } else {
            return nums[p]
        }
    }
    return quickSelect(0, len(nums)-1)
}
