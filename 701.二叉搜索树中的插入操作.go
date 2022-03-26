package main

/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return root
	}

	var insert func(*TreeNode, int)

	insert = func(node *TreeNode, val int) {
		if val > node.Val {
			if node.Right == nil {
				n := &TreeNode{
					Val: val,
				}
				node.Right = n
			} else {
				insert(node.Right, val)
			}
		} else {
			if node.Left == nil {
				n := &TreeNode{
					Val: val,
				}
				node.Left = n

			} else {
				insert(node.Left, val)
			}

		}

	}
	insert(root, val)
	return root

}

// @lc code=end
