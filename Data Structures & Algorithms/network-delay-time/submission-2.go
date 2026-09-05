type Edge struct {
	node, weight int
}

type MinHeap []Edge
func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool  { return h[i].weight < h[j].weight }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) {*h = append(*h, x.(Edge))}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}

func networkDelayTime(times [][]int, n int, k int) int {
    edges := map[int][]Edge{}
	for _, time := range times {
		edges[time[0]] = append(edges[time[0]], Edge{node: time[1], weight: time[2]})
	}

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Edge{node: k, weight: 0})

	visit := map[int]bool{}
	t := 0
	for minHeap.Len() > 0 {
		edge := heap.Pop(minHeap).(Edge)
		if visit[edge.node] {continue}
		visit[edge.node] = true
		t = edge.weight
		for _, next := range edges[edge.node] {
			if visit[next.node] {continue}
			heap.Push(minHeap, Edge{node: next.node, weight: t + next.weight})
		}
	}

	if len(visit) == n {return t}
	return -1
}
