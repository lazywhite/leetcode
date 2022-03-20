/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for s[right] == byte(32) {
		right--
	}
	left := right

	for left >= 0 && s[left] != byte(32)  {
		left--
	}
	return right - left
}

// @lc code=end

