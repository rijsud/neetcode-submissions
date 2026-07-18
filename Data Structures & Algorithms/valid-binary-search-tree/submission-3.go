/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return valid(root, math.MinInt64, math.MaxInt64)
}

func valid(root *TreeNode, left int64, right int64) bool {
	if root == nil {return true}

	val := int64(root.Val)
	if val <= left || val >= right {return false}

	return valid(root.Left, left, val) && valid(root.Right, val, right)
}
