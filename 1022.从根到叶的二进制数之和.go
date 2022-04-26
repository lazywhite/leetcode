/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
 *
 * https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers/description/
 *
 * algorithms
 * Easy (71.27%)
 * Likes:    141
 * Dislikes: 0
 * Total Accepted:    23.6K
 * Total Submissions: 33.1K
 * Testcase Example:  '[1,0,1,0,1,0,1]'
 *
 * 给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
 *
 *
 * 例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。
 *
 *
 * 对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
 *
 * 返回这些数字之和。题目数据保证答案是一个 32 位 整数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,0,1,0,1,0,1]
 * 输出：22
 * 解释：(100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [0]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 树中的节点数在 [1, 1000] 范围内
 * Node.val 仅为 0 或 1
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
func sumRootToLeaf(root *TreeNode) int {

	ans := 0
	var dfs func(*TreeNode, int)
	dfs = func(n *TreeNode, pre int) {
		// 跟pre<<1 + n.Val 值一样
		pre = pre<<1 | n.Val
		if n.Left == nil && n.Right == nil {
			ans += pre
		}
		if n.Left != nil {
			dfs(n.Left, pre)
		}
		if n.Right != nil {
			dfs(n.Right, pre)
		}
	}
	dfs(root, 0)
	return ans

}

// @lc code=end

