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

func networkDelayTime(times [][]int, n int, k int) int {
	hashMap := map[int][][]int{}
	for _, time := range times {
		hashMap[time[0]] = append(hashMap[time[0]], []int{time[1], time[2]})
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, []int{0, k})

	visit := map[int]bool{}
	t := 0

	for minHeap.Len() > 0 {
		path := heap.Pop(minHeap).([]int)
		if visit[path[1]] {continue}
		visit[path[1]] = true
		t = max(t, path[0])
		for _, neighbor := range hashMap[path[1]] {
			if visit[neighbor[0]] {continue}
			heap.Push(minHeap, []int{path[0] + neighbor[1], neighbor[0]})
		}
	}
	
	if len(visit) == n {return t}
	return -1
}
