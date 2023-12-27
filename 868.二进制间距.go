/*
 * @lc app=leetcode.cn id=868 lang=golang
 *
 * [868] 二进制间距
 */

// @lc code=start
func binaryGap(n int) int {
	ans := 0
	i := 0
	last := -1
	for n > 0 {
		if n&1 == 1 {
			if last != -1 {
				ans = max(ans, i-last)
				//ans = max(ans, i-last-1) codility中gap定义不同, 需要减一
			}
			last = i
		}
		n >>= 1
		i++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

