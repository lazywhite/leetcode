/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x != 0 && x%10 == 0 {
		return false
	}
	var res = 0

	for res < x {
		res = res*10 + x%10
		x /= 10
	}

	if res == x || res/10 == x {
		return true
	}
	return false
}

// @lc code=end

