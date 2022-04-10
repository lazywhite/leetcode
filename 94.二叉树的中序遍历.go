/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		result = append(result, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return result

}

// @lc code=end

