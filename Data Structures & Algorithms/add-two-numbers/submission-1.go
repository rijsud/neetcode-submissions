/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
    if l1 == nil && l2 == nil && carry == 0 {
        return nil
    }

    v1, v2 := 0, 0
    if l1 != nil {
        v1 = l1.Val
    }
    if l2 != nil {
        v2 = l2.Val
    }

    sum := v1 + v2 + carry
    carry, val := sum/10, sum%10

    var nextNode *ListNode
    nextL1 := l1
    nextL2 := l2
    if l1 != nil {
        nextL1 = l1.Next
    } else {
        nextL1 = nil
    }
    if l2 != nil {
        nextL2 = l2.Next
    } else {
        nextL2 = nil
    }
    nextNode = add(nextL1, nextL2, carry)

    return &ListNode{Val: val, Next: nextNode}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    return add(l1, l2, 0)
}