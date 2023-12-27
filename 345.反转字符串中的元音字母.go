/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	pt := "aeiouAEIOU"
	t := []byte(s)
	n := len(t)
	i, j := 0, n-1
	for i < j {
		for i < n && !strings.Contains(pt, string(t[i])) {
			i++
		}
		for j > 0 && !strings.Contains(pt, string(t[j])) {
			j--
		}
		if i < j {
			t[i], t[j] = t[j], t[i]
			i++
			j--
		}
	}
	return string(t)
}

// @lc code=end

