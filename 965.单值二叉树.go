/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
 *
 * https://leetcode-cn.com/problems/univalued-binary-tree/description/
 *
 * algorithms
 * Easy (68.65%)
 * Likes:    107
 * Dislikes: 0
 * Total Accepted:    39.5K
 * Total Submissions: 57.6K
 * Testcase Example:  '[1,1,1,1,1,null,1]'
 *
 * 如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
 *
 * 只有给定的树是单值二叉树时，才返回 true；否则返回 false。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：[1,1,1,1,1,null,1]
 * 输出：true
 *
 *
 * 示例 2：
 *
 *
 *
 * 输入：[2,2,2,5,2]
 * 输出：false
 *
 *
 *
 *
 * 提示：
 *
 *
 * 给定树的节点数范围是 [1, 100]。
 * 每个节点的值都是整数，范围为 [0, 99] 。
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
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var dfs func(*TreeNode) bool

	dfs = func(n *TreeNode) bool {
		if n.Left == nil && n.Right == nil {
			return true
		}
		if n.Left != nil && n.Right != nil {
			return n.Val == n.Left.Val && n.Val == n.Right.Val && isUnivalTree(n.Left) && isUnivalTree(n.Right)
		}
		if n.Left != nil {
			return n.Val == n.Left.Val && isUnivalTree(n.Left)
		}
		return n.Val == n.Right.Val && isUnivalTree(n.Right)
	}

	return dfs(root)
}

// @lc code=end

