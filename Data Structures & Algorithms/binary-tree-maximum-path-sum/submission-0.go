/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
    globalMax := math.MinInt

    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return math.MinInt
        }

        left, right := max(0, dfs(node.Left)), max(0, dfs(node.Right))
        disjoinedPath := node.Val + left + right
        continuedPath := node.Val + max(left, right)

        globalMax = max(globalMax, max(disjoinedPath, continuedPath))

        return continuedPath
    }

    dfs(root)
    return globalMax
}