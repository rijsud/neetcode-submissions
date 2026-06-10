/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
	node := dummy
    carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		temp := &ListNode{}
        v1, v2 := 0, 0
        if l1 != nil {
            v1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            v2 = l2.Val
            l2 = l2.Next
        }
		temp.Val = v1 + v2 + carry
        carry = temp.Val / 10
        temp.Val = temp.Val % 10

        node.Next = temp
		node = node.Next
	}

    return dummy.Next
}
