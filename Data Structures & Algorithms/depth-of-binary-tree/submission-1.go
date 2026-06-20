/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type nodeStack struct {
	node *TreeNode
	depth int
}

func maxDepth(root *TreeNode) int {
    stack := []nodeStack{{node: root, depth: 1}}
	res := 0

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp.node != nil {
			res = max(res, temp.depth)
			stack = append(stack, nodeStack{node: temp.node.Left, depth: temp.depth + 1} )
			stack = append(stack, nodeStack{node: temp.node.Right, depth: temp.depth + 1} )
		}
	}

	return res
}
