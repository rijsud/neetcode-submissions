type MinHeap [][]int
func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.([]int))}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	pop := old[n-1]
	*h = old[:n-1]
	return pop
}

func kClosest(points [][]int, k int) [][]int {
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for _, point := range points {
		dist := point[0]*point[0] + point[1]*point[1]
		heap.Push(minHeap, []int{dist, point[0], point[1]})
	}

	res := [][]int{}
	for k > 0 {
		point := heap.Pop(minHeap).([]int)
		res = append(res, []int{point[1], point[2]})
		k--
	}

	return res
}
