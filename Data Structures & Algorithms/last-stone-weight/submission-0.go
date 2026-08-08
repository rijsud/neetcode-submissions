type MaxHeap []int
func (h MaxHeap) Len() int {return len(h)}
func (h MaxHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x interface{}) {*h = append(*h, x.(int))}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	pop := old[n-1]
	*h = old[:n-1]
	return pop
}

func lastStoneWeight(stones []int) int {
	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, stone := range stones {
		heap.Push(maxHeap, stone)
	}
	for maxHeap.Len() > 1 {
		stone1 := heap.Pop(maxHeap).(int)
		stone2 := heap.Pop(maxHeap).(int)
		stone1 = stone1 - stone2
		if stone1 != 0 { heap.Push(maxHeap, stone1) }
	}

	if maxHeap.Len() == 0 {return 0}
	return heap.Pop(maxHeap).(int)
}
