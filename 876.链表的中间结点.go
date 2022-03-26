/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 快慢指针法, fast是slow的2倍, fast到达末尾时, slow刚好走到中间
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil {
		if fast.Next == nil {
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// @lc code=end

