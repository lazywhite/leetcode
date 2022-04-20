/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 *
 * https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 *
 * algorithms
 * Medium (38.57%)
 * Likes:    7338
 * Dislikes: 0
 * Total Accepted:    1.6M
 * Total Submissions: 4.3M
 * Testcase Example:  '"abcabcbb"'
 *
 * 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: s = "abcabcbb"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
 *
 *
 * 示例 2:
 *
 *
 * 输入: s = "bbbbb"
 * 输出: 1
 * 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
 *
 *
 * 示例 3:
 *
 *
 * 输入: s = "pwwkew"
 * 输出: 3
 * 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
 * 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= s.length <= 5 * 10^4
 * s 由英文字母、数字、符号和空格组成
 *
 *
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	m := make(map[byte]int, 0)

	maxLen := 0
	preLen := 0
	for i := range s {
		char := s[i] // byte
		if v, ok := m[char]; ok {
			m[char] = i
			preLen = i - v
			// 删除index小于v的所有key
			cleanCache(m, v)
		} else {
			m[char] = i
			currentLen := preLen + 1
			maxLen = max(currentLen, maxLen)
			preLen = currentLen
		}

	}
	return maxLen
}

func cleanCache(m map[byte]int, i int) {
	for k, v := range m {
		if v <= i {
			delete(m, k)
		}
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

