type Node struct {
	i, j, elevation int
}

type MinHeap []Node

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].elevation < h[j].elevation }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Node)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func swimInWater(grid [][]int) int {
	n := len(grid)
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	t := 0

    minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, Node{i : 0, j : 0, elevation : grid[0][0]})
	visit := map[Node]bool{}

	for minHeap.Len() > 0 {
		node := heap.Pop(minHeap).(Node)
		if visit[node] {continue}
		visit[node] = true
		if t < node.elevation {t = node.elevation}
		if node.i == n - 1 && node.j == n - 1 {return t}
		for _, dir := range dirs {
			nextNodeI := node.i + dir[0]
			nextNodeJ := node.j + dir[1]
			if nextNodeI < 0 || nextNodeJ < 0 || nextNodeI == n || nextNodeJ == n {continue}
			heap.Push(minHeap, Node{i : nextNodeI, j : nextNodeJ, elevation : grid[nextNodeI][nextNodeJ]})
		}
	}

	return -1
}
