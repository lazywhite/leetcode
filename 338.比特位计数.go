/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start
func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = countOnes(i)
	}

	return res
}

// Brian Kernighan 算法
// x & (x - 1) 可以将2进制数字的最后一个1变为0, 直到所有的位都为0
func countOnes(x int) int {
	count := 0
	for x > 0 {
		x &= (x - 1)
		count++
	}
	return count
}

// @lc code=end

