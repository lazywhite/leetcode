/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	// 必须要有, 否则死循环
	if n == 0 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}
	return n == 1
}

// @lc code=end

