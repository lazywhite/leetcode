/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	result := []int{}

	var pre func(*Node)
	pre = func(n *Node) {
		if n == nil {
			return
		}
		result = append(result, n.Val)
		for i := range n.Children {
			pre(n.Children[i])
		}
	}
	pre(root)
	return result

}

// @lc code=end

