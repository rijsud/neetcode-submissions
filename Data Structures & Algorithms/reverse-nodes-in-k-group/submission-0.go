/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy
	for {
		kth := findKth(groupPrev, k)
		if kth == nil {break}
		groupNext := kth.Next

		prev, curr := groupNext, groupPrev.Next
		for curr != groupNext {
			temp := curr.Next
			curr.Next = prev
			prev = curr
			curr = temp
		}

		temp := groupPrev.Next
		groupPrev.Next = prev
		groupPrev = temp
	}
	return dummy.Next
}

func findKth(node *ListNode, k int) *ListNode {
	for k > 0  && node != nil {
		k--
		node = node.Next
	}
	return node
}
