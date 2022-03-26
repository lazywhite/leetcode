/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
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

// 前序遍历, left child不能省略
func tree2str(node *TreeNode) string {
	if node == nil {
		return ""
	}
	if node.Left == nil && node.Right == nil {
		return fmt.Sprintf("%d", node.Val)
	}
	left := tree2str(node.Left)
	right := tree2str(node.Right)
	if right == "" {
		return fmt.Sprintf("%d(%s)", node.Val, left)
	}
	return fmt.Sprintf("%d(%s)(%s)", node.Val, left, right)
}

// @lc code=end

