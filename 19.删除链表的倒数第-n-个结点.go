/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head

	fast := head
	slow := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next

}

// @lc code=end

// 无法处理长度只有1的数组
func oneMoreStep(head *ListNode, n int) *ListNode {

	fast := head
	slow := head

	// 让fast多走1步
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	// fast到末尾时, slow刚好走到-n的父节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}
