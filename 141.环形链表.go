/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Floyd 判圈算法
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil {
		// 因为fast肯定最先达到nil, 因此可以不判断slow
		// fast已经到达了末尾, 说明无环
		if fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}
	return false
}

// @lc code=end

