func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int)
    freqList := make([][]int, len(nums) + 1)

    for _, number := range nums {
        freqMap[number]++
    }

    for number, count := range freqMap {
        freqList[count] = append(freqList[count], number)
    }

    var result []int

    for i := len(freqList) - 1; i > 0; i-- {
        for _, number := range freqList[i] {
            result = append(result, number)
            if len(result) == k {
                return result
            }
        }
    }

    return result
}
