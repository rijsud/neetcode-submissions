/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    res := root.Val
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {return 0}
		maxLeft := dfs(root.Left)
		maxRight := dfs(root.Right)
		maxLeft = max(maxLeft, 0)
		maxRight = max(maxRight, 0)

		res = max(res, root.Val + maxLeft + maxRight)
		return root.Val + max(maxLeft, maxRight)
	}

	dfs(root)
	return res
}
