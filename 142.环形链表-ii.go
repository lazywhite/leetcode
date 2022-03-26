/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	// if head == nil {
	// 	return nil
	// }
	// 起使条件改成了都从head开始
	slow := head
	fast := head

	// 需要修改判定条件, 否则无法进入循环
	// for slow != fast {
	for fast != nil {
		// 先让两个指针移动, 否则直接相等了
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next

		if fast == slow {
			// 相遇之后让ptr从head开始移动, 与slow相遇的地方即为入口
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}

// @lc code=end
