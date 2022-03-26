/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var rt byte

	for i := range s {
		rt = rt ^ s[i]
	}
	for i := range t {
		rt = rt ^ t[i]
	}
	return rt
}

// @lc code=end

