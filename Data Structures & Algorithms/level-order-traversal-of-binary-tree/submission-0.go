/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
	if root == nil {return res}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		queueLen := len(queue)
		list := []int{}

		for i := 0; i < queueLen; i++ {
			node := queue[0]
			queue = queue[1:]

			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, list)
	}

	return res
}
