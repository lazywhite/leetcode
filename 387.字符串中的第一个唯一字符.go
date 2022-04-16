/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
func firstUniqChar(s string) int {
	cache := [26]int{}

	// 第一次遍历s, 缓存所有char次数
	for i := range s {
		cache[s[i]-byte('a')]++
	}
	// 第二次遍历s, 返回第一个cache值为1的index
	for i := range s {
		if cache[s[i]-byte('a')] == 1 {
			return i
		}
	}
	return -1
}

// @lc code=end

