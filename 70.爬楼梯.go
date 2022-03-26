/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	penult := 1
	last := 2
	for i := 3; i <= n; i++ {
		now := penult + last
		penult = last
		last = now
	}
	return last
}

// @lc code=end

