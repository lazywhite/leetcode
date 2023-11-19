/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	c := make([]int, 26)
	for _, char := range s {
		c[char-'a']++
	}
	for _, char := range t {
		c[char-'a']--
	}
	for _, count := range c {
		if count != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

