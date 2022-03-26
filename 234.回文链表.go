/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	res := make([]int, 0)

	current := head
	for current != nil {
		res = append(res, current.Val)
		current = current.Next
	}

	size := len(res)
	if size == 0 {
		return false
	}
	if size == 1 {
		return true
	}

	left := 0
	right := size - 1

	for left <= right {
		if res[left] != res[right] {
			return false
		}
		left++
		right--
	}

	// left > right 仍然没有发现不同
	return true

}

// @lc code=end

