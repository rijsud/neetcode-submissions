/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	newNode := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			newNode.Next = list1
			list1 = list1.Next
		} else {
			newNode.Next = list2
			list2 = list2.Next
		}
		newNode = newNode.Next
	}

	if list1 != nil {
		newNode.Next = list1
	} else if list2 != nil {
		newNode.Next = list2
	}

	return head.Next
}
