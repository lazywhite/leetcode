/*
 * @lc app=leetcode.cn id=1446 lang=golang
 *
 * [1446] 连续字符
 *
 * https://leetcode-cn.com/problems/consecutive-characters/description/
 *
 * algorithms
 * Easy (60.93%)
 * Likes:    98
 * Dislikes: 0
 * Total Accepted:    53K
 * Total Submissions: 86.9K
 * Testcase Example:  '"leetcode"'
 *
 * 给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。
 *
 * 请你返回字符串 s 的 能量。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "leetcode"
 * 输出：2
 * 解释：子字符串 "ee" 长度为 2 ，只包含字符 'e' 。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abbcccddddeeeeedcba"
 * 输出：5
 * 解释：子字符串 "eeeee" 长度为 5 ，只包含字符 'e' 。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 500
 * s 只包含小写英文字母。
 *
 *
 */

// @lc code=start
func maxPower(s string) int {

	max := 1
	pre := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur := pre + 1
			if cur > max {
				max = cur
			}
			pre = cur
		} else {
			pre = 1
		}
	}
	return max
}

// @lc code=end

