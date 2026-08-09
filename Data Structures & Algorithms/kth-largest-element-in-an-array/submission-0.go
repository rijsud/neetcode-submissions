type MinHeap []int
func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    pop := old[n-1]
    *h = old[:n-1]
    return pop
}

func findKthLargest(nums []int, k int) int {
    minHeap := &MinHeap{}
    heap.Init(minHeap)

    for _, num := range nums {
        heap.Push(minHeap, num)
        if minHeap.Len() > k {
            heap.Pop(minHeap)
        }
    }
    return heap.Pop(minHeap).(int)
}
