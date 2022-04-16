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
	if list1.Val < list2.Val {
		// 认为list1为head
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	// 认为list2为head
	list2.Next = mergeTwoLists(list2.Next, list1)
	return list2
}

// @lc code=end

