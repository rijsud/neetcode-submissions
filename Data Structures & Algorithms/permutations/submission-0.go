//recursive

func permute(nums []int) [][]int {
    if len(nums) == 0 {return [][]int{{}}}

    perms := permute(nums[1:])
    res := [][]int{}

    for _, p := range perms {
        for i := 0; i <= len(p); i++ {
            pCopy := make([]int, len(p)+1)
            copy(pCopy[:i], p[:i])
            pCopy[i] = nums[0]
            copy(pCopy[i+1:], p[i:])
            res = append(res, pCopy)
        }
    }
    return res
}
