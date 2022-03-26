/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
func toLowerCase(s string) string {

	bytes := []byte(s)

	for i := range bytes {
		// Z: 90
		if bytes[i] >= 65 && bytes[i] <= 90 {
			bytes[i] += 32
		}
	}
	return string(bytes)

}

// @lc code=end

