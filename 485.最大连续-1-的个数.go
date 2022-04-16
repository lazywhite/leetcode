/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 *
 * https://leetcode-cn.com/problems/max-consecutive-ones/description/
 *
 * algorithms
 * Easy (60.95%)
 * Likes:    313
 * Dislikes: 0
 * Total Accepted:    148.8K
 * Total Submissions: 244K
 * Testcase Example:  '[1,1,0,1,1,1]'
 *
 * 给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：nums = [1,1,0,1,1,1]
 * 输出：3
 * 解释：开头的两位和最后的三位都是连续 1 ，所以最大连续 1 的个数是 3.
 *
 *
 * 示例 2:
 *
 *
 * 输入：nums = [1,0,1,1,0,1]
 * 输出：2
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * nums[i] 不是 0 就是 1.
 *
 *
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {

	maxLen := 0
	preLen := 0

	for i := range nums {
		num := nums[i]

		if num == 1 {
			curLen := preLen + 1
			maxLen = max(maxLen, curLen)
			preLen = curLen
		} else {
			preLen = 0
		}
	}
	return maxLen
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

