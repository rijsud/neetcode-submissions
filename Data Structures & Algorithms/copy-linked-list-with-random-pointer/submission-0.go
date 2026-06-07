/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
    oldToCopy := map[*Node]*Node{nil: nil}
	cur := head
	for cur != nil {
		newNode := &Node{Val:cur.Val}
		oldToCopy[cur] = newNode
		cur = cur.Next
	}

	cur = head
	for cur != nil {
		newNode := oldToCopy[cur]
		newNode.Next = oldToCopy[cur.Next]
		newNode.Random = oldToCopy[cur.Random]
		cur = cur.Next
	}

	return oldToCopy[head]
}
