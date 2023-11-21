/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	/* 分治 + 合并链表 */
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	start := 0
	end := len(lists) - 1

	// 合并两个有序链表
	var mergeTwo func(*ListNode, *ListNode) *ListNode
	mergeTwo = func(a, b *ListNode) *ListNode {
		if a == nil {
			return b
		}
		if b == nil {
			return a
		}
		/*
			不能使用递归, 否则报内存溢出
			if a.Val < b.Val {
				a.Next = mergeTwo(a.Next, b)
				return a
			} else {
				b.Next = mergeTwo(b.Next, b)
				return b
			}
		*/
		head := &ListNode{
			Val: 0,
		}
		tail := head
		for a != nil && b != nil {
			if a.Val < b.Val {
				tail.Next = a
				a = a.Next
			} else {
				tail.Next = b
				b = b.Next
			}
			tail = tail.Next
		}
		if a == nil {
			tail.Next = b
		}
		if b == nil {
			tail.Next = a
		}
		return head.Next
	}

	// 分治, 二等拆分
	var split func(start, end int) *ListNode
	split = func(start, end int) *ListNode {
		// 拆到只剩一个元素
		if start == end {
			return lists[start]
		}
		mid := start + (end-start)/2
		left := split(start, mid)
		right := split(mid+1, end)
		return mergeTwo(left, right)
	}
	return split(start, end)
}

// @lc code=end

