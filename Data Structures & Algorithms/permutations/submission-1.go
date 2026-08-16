//iterative

func permute(nums []int) [][]int {
    perms := [][]int{{}}

    for _, num := range nums {
        newPerms := [][]int{}
        for _, p := range perms {
            for i := 0; i <= len(p); i++ {
                pCopy := make([]int, len(p)+1)
                copy(pCopy[:i], p[:i])
                pCopy[i] = num
                copy(pCopy[i+1:], p[i:])
                newPerms = append(newPerms, pCopy)
            }
        }
        perms = newPerms
    }
    return perms
}
