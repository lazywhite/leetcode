/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {

	depth := 0
	if root == nil {
		return depth
	}
	queue := []*Node{
		root,
	}

	for len(queue) > 0 {
		depth++
		size := len(queue)
		for i := 0; i < size; i++ {
			queue = append(queue, queue[i].Children...)
		}
		queue = queue[size:]
	}
	return depth
}

// @lc code=end

