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

	var dfs func(preorder []int, inorder []int) *TreeNode
	dfs = func(preorder []int, inorder []int) *TreeNode {
		if len(preorder) == 0 || len(inorder) == 0 {return nil}
		root := &TreeNode{Val: preorder[0]}
		mid := indices[preorder[0]]
		root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
		root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
		return root
	}

	return dfs(preorder, inorder)
}