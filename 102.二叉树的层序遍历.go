/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n1 := &TreeNode{
		Val: 3,
	}
	n2 := &TreeNode{
		Val: 9,
	}
	n3 := &TreeNode{
		Val: 20,
	}
	n4 := &TreeNode{
		Val: 15,
	}
	n5 := &TreeNode{
		Val: 7,
	}

	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	x := levelOrder(n1)
	fmt.Println(x)
}


// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)
		levelData := make([]int, 0)
		for i := 0; i < size; i++ {
			// get node
			node := queue[i]

			// append value
			levelData = append(levelData, node.Val)

			// put children into queue
			// 做nil判断, 否则nil也会被放进queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 最后整体pop
		queue = queue[size:]
		if len(levelData) > 0 {
			result = append(result, levelData)
		}
	}
	return result
}

// @lc code=end
