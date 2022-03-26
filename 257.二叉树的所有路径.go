package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
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
func binaryTreePaths(root *TreeNode) []string {

	result := []string{}

	var path func(*TreeNode, string)
	path = func(n *TreeNode, prefix string) {
		if n == nil {
			return
		}
		current := fmt.Sprintf("%s%d->", prefix, n.Val)
		// 到达leaf node
		if n.Left == nil && n.Right == nil {
			result = append(result, current[:len(current)-2])
		}
		if n.Left != nil {
			path(n.Left, current)
		}
		if n.Right != nil {
			path(n.Right, current)
		}

	}
	path(root, "")
	return result

}

// @lc code=end
