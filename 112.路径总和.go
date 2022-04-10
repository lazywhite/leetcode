/*
 * @lc app=leetcode.cn id=112 lang=golang
 *
 * [112] 路径总和
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
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var has func(*TreeNode, int) bool
	has = func(node *TreeNode, parentSum int) bool {
		curVal := parentSum + node.Val

		if node.Left == nil && node.Right == nil {
			return curVal == targetSum
		}

		leftHas, rightHas := false, false
		if node.Left != nil {
			leftHas = has(node.Left, curVal)
		}
		if node.Right != nil {
			rightHas = has(node.Right, curVal)
		}
		return leftHas || rightHas
	}
	return has(root, 0)
}

// @lc code=end

