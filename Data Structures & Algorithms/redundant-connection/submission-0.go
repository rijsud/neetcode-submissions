func findRedundantConnection(edges [][]int) []int {
	n := len(edges)
	parent, rank := make([]int, n + 1), make([]int, n + 1)
	for i := range parent {
		parent[i] = i
		rank[i] = 1
	}

	find := func(node int) int {
		res := node
		for res != parent[res] {
			parent[res] = parent[parent[res]]
			res = parent[res]
		}
		return res
	}

	union := func(n1, n2 int) bool {
		p1, p2 := find(n1), find(n2)
		if p1 == p2 {return true}
		if rank[p2] > rank[p1] {
			p2, p1 = p1, p2
		}
		parent[p2] = p1
		rank[p1] += rank[p2]
		return false
	}

	for _, edge := range edges {
		if union(edge[0], edge[1]) {return edge}
	}

	return []int{}
}
