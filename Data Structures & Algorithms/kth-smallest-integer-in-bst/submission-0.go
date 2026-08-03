/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    list := []int{}
	res := 0
	if root == nil {return res}
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {return}
		dfs(root.Left)
		list = append(list, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return list[k-1]
}
