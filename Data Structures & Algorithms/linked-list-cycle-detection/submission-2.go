/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    tort, hare := head, head
    for hare != nil && hare.Next != nil {
        tort = tort.Next
        hare = hare.Next.Next
        if tort == hare {
            return true
        }
    }
    return false
}
