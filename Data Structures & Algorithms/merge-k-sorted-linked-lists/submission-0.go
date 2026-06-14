/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    nodes := make([]int, 0)

    for _, list := range lists {
        curr := list
        for curr != nil {
            nodes = append(nodes, curr.Val)
            curr = curr.Next
        }
    }

    sort.Ints(nodes)

    dummy := &ListNode{Val: 0}
    curr := dummy

    for _, val := range nodes {
        curr.Next = &ListNode{Val: val}
        curr = curr.Next
    }

    return dummy.Next
}