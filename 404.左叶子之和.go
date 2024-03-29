/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (59.83%)
 * Likes:    435
 * Dislikes: 0
 * Total Accepted:    142.9K
 * Total Submissions: 237.2K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定二叉树的根节点 root ，返回所有左叶子之和。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 *
 * 输入: root = [3,9,20,null,null,15,7]
 * 输出: 24
 * 解释: 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 *
 *
 * 示例 2:
 *
 *
 * 输入: root = [1]
 * 输出: 0
 *
 *
 *
 *
 * 提示:
 *
 *
 * 节点数在 [1, 1000] 范围内
 * -1000 <= Node.val <= 1000
 *
 *
 *
 *
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

func isLeaf(n *TreeNode) bool {
	if n.Left == nil && n.Right == nil {
		return true
	}
	return false
}

func sumOfLeftLeaves(root *TreeNode) int {

	result := 0
	var dfs func(*TreeNode)

	dfs = func(n *TreeNode) {
		if isLeaf(n) {
			return
		}
		if n.Left != nil {
			if isLeaf(n.Left) {
				result += n.Left.Val
			} else {
				dfs(n.Left)
			}
		}
		if n.Right != nil {
			if !isLeaf(n.Right) {
				dfs(n.Right)
			}
		}
	}

	dfs(root)
	return result
}

// @lc code=end

