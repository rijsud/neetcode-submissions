func topKFrequent(nums []int, k int) []int {
    freqMap := make(map[int]int) // map of number with frequency
    for _, number := range nums {
        freqMap[number]++
    }

    freqList := make([][2]int, 0, len(freqMap))

    for num, cnt := range freqMap {
        freqList = append(freqList, [2]int{cnt, num})
    }

    sort.Slice(freqList, func(countFirst, countSecond int) bool {
        return freqList[countFirst][0] > freqList[countSecond][0]
    })

    result := make([]int, k)

    for i := 0; i < k; i++ {
        result[i] = freqList[i][1]
    }

    return result
}
