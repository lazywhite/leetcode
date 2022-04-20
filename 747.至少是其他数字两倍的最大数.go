/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 *
 * https://leetcode-cn.com/problems/largest-number-at-least-twice-of-others/description/
 *
 * algorithms
 * Easy (45.79%)
 * Likes:    170
 * Dislikes: 0
 * Total Accepted:    82.1K
 * Total Submissions: 179.2K
 * Testcase Example:  '[3,6,1,0]'
 *
 * 给你一个整数数组 nums ，其中总是存在 唯一的 一个最大整数 。
 *
 * 请你找出数组中的最大元素并检查它是否 至少是数组中每个其他数字的两倍 。如果是，则返回 最大元素的下标 ，否则返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [3,6,1,0]
 * 输出：1
 * 解释：6 是最大的整数，对于数组中的其他整数，6 至少是数组中其他元素的两倍。6 的下标是 1 ，所以返回 1 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：nums = [1,2,3,4]
 * 输出：-1
 * 解释：4 没有超过 3 的两倍大，所以返回 -1 。
 *
 * 示例 3：
 *
 *
 * 输入：nums = [1]
 * 输出：0
 * 解释：因为不存在其他数字，所以认为现有数字 1 至少是其他数字的两倍。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 50
 * 0 <= nums[i] <= 100
 * nums 中的最大元素是唯一的
 *
 *
 */

// @lc code=start
func dominantIndex(nums []int) int {
	m1, m2, idx := -1, -1, 0
	for i, num := range nums {
		if num > m1 {
			m1, m2, idx = num, m1, i
		} else if num > m2 {
			m2 = num
		}
	}
	if m1 >= m2*2 {
		return idx
	}
	return -1
}

// @lc code=end

