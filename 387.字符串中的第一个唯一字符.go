/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	// bitmap
	cache := [26]int{}

	for i := range s {
		cache[s[i]-byte('a')]++
	}
	for i := range s {
		if cache[s[i]-byte('a')] == 1 {
			return i
		}
	}
	return -1

}

// @lc code=end

