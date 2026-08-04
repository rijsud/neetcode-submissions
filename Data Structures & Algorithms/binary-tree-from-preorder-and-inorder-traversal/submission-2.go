/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 // O(n) time

func buildTree(preorder []int, inorder []int) *TreeNode {
	indices := map[int]int{}
	for idx, val := range inorder {
		indices[val] = idx
	}

	preIdx := 0

	var dfs func(int, int) *TreeNode
	dfs = func(left int, right int) *TreeNode {
		if left > right {return nil}

		root := &TreeNode{Val: preorder[preIdx]}
		mid := indices[preorder[preIdx]]
		preIdx++
		
		root.Left = dfs(left, mid - 1)
		root.Right = dfs(mid + 1, right)
		return root
	}

	return dfs(0, len(preorder)-1)
}