/*
 * @lc app=leetcode.cn id=108 lang=golang
 *
 * [108] 将有序数组转换为二叉搜索树
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
func sortedArrayToBST(nums []int) *TreeNode {
	/*
		bst的inorder结果即为递增数组
		因此每次选择下列节点为根节点均可
		 1. 中间左边
		 2. 中间右边
		 3. 中间
	*/

	return bst(nums, 0, len(nums)-1)
}

func bst(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	root := new(TreeNode)

	// left <= right
	mid := (left + right) / 2
	root.Val = nums[mid]
	root.Left = bst(nums, left, mid-1)
	root.Right = bst(nums, mid+1, right)
	return root
}

// @lc code=end

