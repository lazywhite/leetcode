package main

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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

// morris
/*
1. 如果cur无左孩子，cur向右移动（cur=cur.right）
2. 如果cur有左孩子，找到cur左子树上最右的节点，记为mostright
	如果mostright的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
	如果mostright的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）
*/
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	var (
		p1           = root
		p2 *TreeNode = nil
	)

	for p1 != nil {
		p2 = p1.Left
		// has left child
		if p2 != nil {
			// found right most node, 不能存在映射
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				// ** 建立映射时需要存储val
				result = append(result, p1.Val)
				// 保存映射
				p2.Right = p1
				// p1移动向left
				p1 = p1.Left
				// 重新开始循环
				// continue
			} else {
				// 又访问到了已存在映射的node
				// 1. 清除掉映射
				p2.Right = nil
				// 2. move to right
				p1 = p1.Right
			}

		} else {
			// ** 没有left child需要存储val
			result = append(result, p1.Val)
			// 没有left child, move to right
			p1 = p1.Right
		}

	}
	return result
}

// @lc code=end

// 递归
func preorderRecursive(root *TreeNode) []int {
	result := make([]int, 0)

	var preorder func(*TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}

	preorder(root)
	return result
}

// stack
func preorderStack(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		// pop stack head
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)

		// put node children
		// 前序遍历, 先放入right child
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}

	}

	return result
}
