/*
 * @lc app=leetcode.cn id=226 lang=golang
 *
 * [226] 翻转二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {

	var invert func(*TreeNode)
	invert = func(n *TreeNode) {
		if n == nil {
			return
		}
		n.Left, n.Right = n.Right, n.Left
		invert(n.Left)
		invert(n.Right)

	}
	invert(root)
	return root
}

// @lc code=end

