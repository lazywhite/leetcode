/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	middle := (len(s) - 1) / 2

	left := 0
	right := len(s) - 1
	for i := 0; i <= middle; i++ {
		if left < right {
			swap(s, left, right)
			left++
			right--
		}
	}
}
func swap(s []byte, x, y int) {
	s[x], s[y] = s[y], s[x]
}

// @lc code=end

