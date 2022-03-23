/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
双指针法

	1. 不相交, 则两个指针在遍历完m+n个node后均为nil而相等, 结束循环并返回nil
	2. 相交
		相交部分为c
		headA: 不相交部分为a
		headB: 不相交部分为b

		如果a==b, 则两个指针走完a之后相等
		如果a!=b, 则第一个指针走完a+c+b, 第二个指针走完b+c+a之后相等

		都会在第一个交点相遇, 此时停止循环并返回node

*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB

	//最差的情况, 两个指针都走了m+n个node,
	// 相交部分为c
	// 1. 如果不相交, 两者均为nil
	// 2. 如果相交, 两者会在第一个交点相遇
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// @lc code=end

