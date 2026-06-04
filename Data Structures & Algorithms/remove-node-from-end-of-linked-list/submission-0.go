/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nodes := []*ListNode{}
    cur := head
    for cur != nil {
        nodes = append(nodes, cur)
        cur = cur.Next
    }

    removeIndex := len(nodes) - n
    if removeIndex == 0 {
        return head.Next
    }

    nodes[removeIndex-1].Next = nodes[removeIndex].Next
    return head
}