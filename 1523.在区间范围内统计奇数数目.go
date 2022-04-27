/*
 * @lc app=leetcode.cn id=1523 lang=golang
 *
 * [1523] 在区间范围内统计奇数数目
 *
 * https://leetcode-cn.com/problems/count-odd-numbers-in-an-interval-range/description/
 *
 * algorithms
 * Easy (51.34%)
 * Likes:    49
 * Dislikes: 0
 * Total Accepted:    24K
 * Total Submissions: 46.8K
 * Testcase Example:  '3\n7'
 *
 * 给你两个非负整数 low 和 high 。请你返回 low 和 high 之间（包括二者）奇数的数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：low = 3, high = 7
 * 输出：3
 * 解释：3 到 7 之间奇数数字为 [3,5,7] 。
 *
 * 示例 2：
 *
 * 输入：low = 8, high = 10
 * 输出：1
 * 解释：8 到 10 之间奇数数字为 [9] 。
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= low <= high <= 10^9
 *
 *
 */

// @lc code=start
func countOdds(low int, high int) int {

	// [0, x]区间内的奇数为 (x+1)/2
	// [a, b]区间内的奇数为 [0, b] - [0, a-1]
	odds := func(n int) int {
		return (n + 1) >> 1
	}

	return odds(high) - odds(low-1)

}

// @lc code=end

