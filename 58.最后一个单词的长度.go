/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	end := len(s) - 1

	for end >= 0 {
		if s[end] != byte(' ') {
			break
		}
		end--
	}

	start := end

	for start >= 0 {
		if s[start] == byte(' ') {
			break
		}
		start--
	}

	return end - start

}

// @lc code=end

