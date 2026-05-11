func threeSum(nums []int) [][]int {
    var res [][]int
    indices := make(map[[3]int]struct{})
    sort.Ints(nums)

    for i := 0; i < len(nums); i++{
        for j := i + 1; j < len(nums); j++ {
            for k := j + 1; k < len(nums); k++ {
                sum := nums[i] + nums[j] + nums[k]
                if sum == 0 {
                    indices[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
                    // res = append(res, []int{nums[i], nums[j], nums[k]})
                }
            }
        }
    }

    for list := range indices {
        res = append(res, []int{list[0], list[1], list[2]})
    }

    return res
}
