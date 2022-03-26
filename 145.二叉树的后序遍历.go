/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) []int {

	result := []int{}
	var postOrder func(*TreeNode)
	postOrder = func(n *TreeNode) {
		if n == nil {
			return
		}
		postOrder(n.Left)
		postOrder(n.Right)
		result = append(result, n.Val)
	}
	postOrder(root)
	return result
}

// @lc code=end

