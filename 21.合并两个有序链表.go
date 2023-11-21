/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 返回merge后的head
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	/*
		递归, 时间复杂度较高
		if list1.Val < list2.Val {
			// 认为list1为head
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		}
		// 认为list2为head
		list2.Next = mergeTwoLists(list2.Next, list1)
	*/

	// 生成fake head
	head := &ListNode{
		Val: 0,
	}
	// tail初始值为head
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	/* 其中一个已到末尾 */
	if list1 == nil {
		tail.Next = list2
	}
	if list2 == nil {
		tail.Next = list1
	}

	return head.Next
}

// @lc code=end

