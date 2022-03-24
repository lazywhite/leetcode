/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	result := []int{}

	var post func(*Node)
	post = func(n *Node) {
		if n == nil {
			return
		}
		for i := range n.Children {
			post(n.Children[i])
		}
		result = append(result, n.Val)
	}
	post(root)
	return result

}

// @lc code=end

