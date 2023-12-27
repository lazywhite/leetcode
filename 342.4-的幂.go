/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	// 4的幂, 二进制奇数位都是1, 偶数位都是0, 因此构造一个奇数位都是0, 偶数位都是1的mask
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}

// @lc code=end

