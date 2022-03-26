/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3 的幂
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	/*
		循环
		if n == 0 {
			return false
		}

		for n%3 == 0 {
			n /= 3
		}
		return n == 1
	*/

	// 32位整数, 是3的幂的最大的数是1162261467, 能整除它即为幂数
	if n <= 0{
		return false
	}
	return 1162261467 % n == 0

}

// @lc code=end

