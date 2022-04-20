/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (53.29%)
 * Likes:    860
 * Dislikes: 0
 * Total Accepted:    241.7K
 * Total Submissions: 453.6K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,3,4,4,5]
 * 输出：[1,2,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1,1,1,2,3]
 * 输出：[2,3]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中节点数目在范围 [0, 300] 内
 * -100 <= Node.val <= 100
 * 题目数据保证链表已经按升序 排列
 *
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head

	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Next.Val == current.Next.Val {
			// 找到重复的值
			v := current.Next.Val
			// 从current开始遍历
			for current.Next != nil && current.Next.Val == v {
				current.Next = current.Next.Next
			}

		} else {
			// 必须不存在重复值, 才能向下迭代
			current = current.Next
		}
	}
	return dummy.Next
}

// @lc code=end

