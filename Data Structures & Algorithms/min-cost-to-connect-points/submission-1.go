type Node struct {
	index, weight int
}

type MinHeap []Node

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].weight < h[j].weight }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Node)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func minCostConnectPoints(points [][]int) int {
	n := len(points)
	adj := make(map[int][]Node)
	for i := 0; i < n; i++ {
		x1, y1 := points[i][0], points[i][1]
		for j := i + 1; j < n; j++ {
			x2, y2 := points[j][0], points[j][1]
			dist := int(math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2)))
			adj[i] = append(adj[i], Node{weight : dist, index : j})
			adj[j] = append(adj[j], Node{weight : dist, index : i})
		}
	}

    visit := map[int]bool{}
	cost := 0
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Node{index : 0, weight : 0})

	for len(visit) < n {
		node := heap.Pop(minHeap).(Node)
		if visit[node.index] {continue}
		cost += node.weight
		visit[node.index] = true
		for _, next := range adj[node.index] {
			if visit[next.index] {continue}
			heap.Push(minHeap, next)
		}
	}

	return cost
}
