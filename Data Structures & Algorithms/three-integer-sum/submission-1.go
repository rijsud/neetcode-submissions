func threeSum(nums []int) [][]int {
    var res [][]int
    sort.Ints(nums)

    for i := 0; i < len(nums); i++ {
        // firstNum := nums[i]
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums [i-1] {
            continue
        }

        j, k := i + 1, len(nums) - 1

        for j < k {
            target := nums[i] + nums[j] + nums[k]
            if target > 0 {
                k--
            }
            if target < 0 {
                j++
            }
            if target == 0 {
                res = append(res, []int{nums[i], nums[j], nums[k]})
                j++
                k--
                for j < k && nums[j] == nums[j-1] {
                    j++
                }
            }
        }
    }

    return res
}
