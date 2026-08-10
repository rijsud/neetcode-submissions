type MaxHeap []int
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {return h[i] > h[j]}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int))}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	pop := old[n-1]
	*h = old[:n-1]
	return pop
}

func leastInterval(tasks []byte, n int) int {
	hashMap := map[byte]int{}
	for _, task := range tasks {
		hashMap[task]++
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for _, freq := range hashMap {
		heap.Push(maxHeap, freq)
	}

	time := 0
	queue := [][]int{}
	for maxHeap.Len() > 0 || len(queue) > 0 {
		time++

		if maxHeap.Len() > 0 {
			taskCount := heap.Pop(maxHeap).(int)-1
			if taskCount > 0 {queue = append(queue, []int{taskCount, time+n})}
		}

		if len(queue) > 0 && queue[0][1] == time {
			heap.Push(maxHeap, queue[0][0])
			queue = queue[1:]
		}
	}

	return time
}
