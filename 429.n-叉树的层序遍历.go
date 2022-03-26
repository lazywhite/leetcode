/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {

	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*Node{
		root,
	}

	for len(queue) > 0 {
		size := len(queue)
		data := []int{}

		for _, node := range queue {
			data = append(data, node.Val)
			if len(node.Children) > 0 {
				queue = append(queue, node.Children...)
			}
		}
		queue = queue[size:]
		result = append(result, data)
	}
	return result

}

// @lc code=end

