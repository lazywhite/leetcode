/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		levelData := make([]int, 0)
		for i := 0; i < size; i++ {
			// get node
			node := queue[i]
			if node == nil {
				continue
			}
			// append value
			levelData = append(levelData, node.Val)
			// put children into queue
			queue = append(queue, node.Left, node.Right)
		}
		// pop node
		queue = queue[size:]
		if len(levelData) > 0 {
			result = append(result, levelData)
		}
	}
	return result

}

// @lc code=end

