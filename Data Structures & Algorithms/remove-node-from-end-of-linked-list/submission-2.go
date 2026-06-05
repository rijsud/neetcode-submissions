/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
    node := head
    for node != nil {
        node = node.Next
        length++
    }

    if length == n {
        return head.Next
    }

    node = head
    for node != nil {
        length--
        if length == n {
            node.Next = node.Next.Next
            break
        }
        node = node.Next
    }
    return head
}
