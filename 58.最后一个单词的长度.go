/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	right := len(s) - 1
	for s[right] == ' ' {
		right--
	}
	left := right

	for left >= 0 && s[left] != ' ' {
		left--
	}
	return right - left
}

// @lc code=end

