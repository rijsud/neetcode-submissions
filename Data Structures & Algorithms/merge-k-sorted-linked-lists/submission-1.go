/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {return nil}

	for len(lists) > 1 {
		mergedList := []*ListNode{}
		for i := 0; i < len(lists); i += 2 {
			l1 := lists[i]
			var l2 *ListNode
			if (i + 1) < len(lists) {l2 = lists[i+1]}
			mergedList = append(mergedList, mergeTwoLists(l1, l2))
		}
		lists = mergedList
	}

	return lists[0]
}

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
