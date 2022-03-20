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
	return has(root, targetSum, 0)
}

func has(node *TreeNode, targetSum, parentSum int) bool {
	curVal := parentSum + node.Val

	if node.Left == nil && node.Right == nil {
		return curVal == targetSum
	}

	leftHas, rightHas := false, false
	if node.Left != nil {
		leftHas = has(node.Left, targetSum, curVal)
	}
	if node.Right != nil {
		rightHas = has(node.Right, targetSum, curVal)
	}
	return leftHas || rightHas
}

// @lc code=end

