/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	oldToNew := map[*Node]*Node{}

    var dfs func(*Node) *Node
	dfs = func(argNode *Node) *Node {
		if argNode == nil {return nil}
		if returnNode, ok := oldToNew[argNode]; ok {
			return returnNode
		}

		copyNode := &Node{Val : argNode.Val}
		oldToNew[argNode] = copyNode
		for _, neighbor := range argNode.Neighbors {
			copyNode.Neighbors = append(copyNode.Neighbors, dfs(neighbor))
		}
		return copyNode
	}

	return dfs(node)
}
