/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	count := 0

	for i := 0; i < 32; i++ {
		// 第i位为1, 结果才大于0
		if 1<<i&num > 0 {
			count++
		}
	}
	return count
}

// @lc code=end

