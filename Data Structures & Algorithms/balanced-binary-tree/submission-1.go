/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type Result struct {
    balanced bool
    height   int
}

func isBalanced(root *TreeNode) bool {
    var dfs func(*TreeNode) Result
    dfs = func(root *TreeNode) Result {
        if root == nil {
            return Result{true, 0}
        }
        left, right := dfs(root.Left), dfs(root.Right)
        balance := left.balanced && right.balanced && abs(left.height - right.height) <= 1

        return Result{balance, 1 + max(left.height, right.height)}
    }

    return dfs(root).balanced
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}