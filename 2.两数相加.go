/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		head, tail *ListNode
		carry      = 0
	)

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			// 移动l1
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			// 移动l2
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		left := sum % 10
		carry = sum / 10

		if head == nil {
			// 初始化tail, head
			head = &ListNode{
				Val: left,
			}
			tail = head
		} else {
			node := &ListNode{
				Val: left,
			}
			// 移动tail
			tail.Next = node
			tail = tail.Next
		}
	}

	if carry > 0 {
		node := &ListNode{
			Val: carry,
		}
		tail.Next = node

	}
	return head

}

// @lc code=end

