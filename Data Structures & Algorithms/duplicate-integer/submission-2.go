func hasDuplicate(nums []int) bool {
    duplicateMap := make(map[int]int)
    for _, number := range nums {
        if _, ok := duplicateMap[number] ; ok {
            return true
        }
        duplicateMap[number] = 1
    }
    return false
}
