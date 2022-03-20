/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start

func deleteDuplicates(head *ListNode) *ListNode {

	current := head
	for current != nil {
		for current.Next != nil && current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}
		current = current.Next
	}
	return head
}

// @lc code=end

