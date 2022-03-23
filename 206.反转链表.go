package main

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 最后一个node是newHead, 所有的递归函数都返回同一个
	newHead := reverseList(head.Next)

	// 上层函数保存了老的映射
	// 由上层函数进行reverse
	head.Next.Next = head
	// 将自身的Next置为nil
	head.Next = nil

	return newHead
}

// @lc code=end
