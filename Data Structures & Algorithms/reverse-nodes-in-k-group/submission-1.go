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

		lastInReversed := groupPrev.Next
		groupPrev.Next = kth
		groupPrev = lastInReversed
	}

	return dummy.Next
}

func findKth(node *ListNode, k int) *ListNode {
	for node != nil && k > 0 {
		node = node.Next
		k--
	}
	return node
}
