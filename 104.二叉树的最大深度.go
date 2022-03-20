/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDep(root, 1)
}

func maxDep(node *TreeNode, depth int) int {

	leftMax := depth
	rightMax := depth
	if node.Left != nil {
		leftMax = maxDep(node.Left, depth+1)
	}
	if node.Right != nil {
		rightMax = maxDep(node.Right, depth+1)
	}
	return max(leftMax, rightMax)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

