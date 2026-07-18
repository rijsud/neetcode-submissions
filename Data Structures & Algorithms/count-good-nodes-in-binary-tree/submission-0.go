/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
    goodNodes := 0
	if root == nil {return goodNodes}
	var dfs func(root *TreeNode, greatest int)
	dfs = func(root *TreeNode, greatest int) {
		if root == nil {return}
		nextGreatest := greatest
		if root.Val >= greatest {
			nextGreatest = root.Val
			goodNodes++
		}
		dfs(root.Left, nextGreatest)
		dfs(root.Right, nextGreatest)
	}

	dfs(root, root.Val)
	return goodNodes
}
