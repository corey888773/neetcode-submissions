/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var lowestCommonAncestor *TreeNode

    var dfs func(node *TreeNode) (bool, bool)
    dfs = func(node *TreeNode) (bool, bool) {
        if node == nil {
            return false, false
        }

        leftP, leftQ := dfs(node.Left)
        rightP, rightQ := dfs(node.Right)
        currP, currQ := node.Val == p.Val, node.Val == q.Val

        foundP := leftP || rightP || currP
        foundQ := leftQ || rightQ || currQ

        if lowestCommonAncestor == nil && foundP && foundQ {
			lowestCommonAncestor = node
		}

        return foundP, foundQ
    }

	dfs(root)
    return lowestCommonAncestor
}