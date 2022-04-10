/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
// KMP
func strStr(target string, pattern string) int {

	n, m := len(target), len(pattern)
	if m == 0 {
		return 0
	}
	pi := make([]int, m)
	// compile pattern
	for i, j := 1, 0; i < m; i++ {
		for j > 0 && pattern[i] != pattern[j] {
			j = pi[j-1]
			// j-- 必须从pi中取值
		}
		if pattern[i] == pattern[j] {
			j++
		}
		pi[i] = j
	}
	for i, j := 0, 0; i < n; i++ {
		for j > 0 && target[i] != pattern[j] {
			j = pi[j-1]
		}
		if target[i] == pattern[j] {
			j++
		}
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

// @lc code=end

