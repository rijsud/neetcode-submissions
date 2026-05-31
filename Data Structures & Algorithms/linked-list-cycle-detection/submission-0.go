/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    seen := make(map[*ListNode]bool)
    cur := head
    for cur != nil {
        if seen[cur] {
            return true
        }
        seen[cur] = true
        cur = cur.Next
    }
    return false
}
