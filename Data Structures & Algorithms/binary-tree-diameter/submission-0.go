/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0

	var dfs func(*TreeNode) int
	dfs = func(curr *TreeNode) int {
		if curr == nil {
			return 0
		}
		left := dfs(curr.Left)
		right := dfs(curr.Right)

		res = max(res, left + right)
		return 1 + max(left, right)
	}

	dfs(root)
	return res
}
