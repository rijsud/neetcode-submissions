/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	hashMap := map[*TreeNode][]int{}
	hashMap[nil] = []int{0,0}

	for len(stack) > 0 {
		node := stack[len(stack)-1]

		if node.Left != nil && len(hashMap[node.Left]) == 0 {
			stack = append(stack, node.Left)
		} else if node.Right != nil && len(hashMap[node.Right]) == 0 {
			stack = append(stack, node.Right)
		} else {
			stack = stack[:len(stack)-1]
			leftHeight := hashMap[node.Left][0]
           leftDiameter := hashMap[node.Left][1]
           rightHeight := hashMap[node.Right][0]
           rightDiameter := hashMap[node.Right][1]

           height := 1 + max(leftHeight, rightHeight)
           diameter := max(leftHeight+rightHeight, max(leftDiameter, rightDiameter))

           hashMap[node] = []int{height, diameter}
		}
	}

	return hashMap[root][1]
}
