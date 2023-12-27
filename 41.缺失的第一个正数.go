/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	n := len(nums)
	m := make([]int, n) // 结果一定是[1, N]间的整数
	for _, e := range nums {
		if e > 0 && e <= n { // 直接忽略不在[1,N]间的数即可
			m[e-1] = 1
		}
	}
	for i, v := range m {
		if v == 0 {
			return i + 1
		}
	}
	return n + 1
}

// @lc code=end

