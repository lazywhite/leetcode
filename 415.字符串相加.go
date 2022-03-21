/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 */

// @lc code=start
func addStrings(num1 string, num2 string) string {

	result := []string{}
	add := 0
	len1 := len(num1)
	len2 := len(num2)
	maxLen := maxLen(len1, len2)

	for i := 0; i < maxLen; i++ {
		var (
			char1 byte = '0'
			char2 byte = '0'
		)

		if i <= len1-1 {
			char1 = num1[len1-1-i]
		}
		if i <= len2-1 {
			char2 = num2[len2-1-i]
		}

		// '0'等于48
		n1 := int(char1) - 48
		n2 := int(char2) - 48
		sum := n1 + n2 + add

		add = sum / 10
		left := sum % 10

		s := strconv.Itoa(left)
		result = append([]string{s}, result...)
	}
	if add == 1 {
		result = append([]string{"1"}, result...)
	}
	final := ""
	for i := range result {
		final += result[i]
	}
	return final

}

func maxLen(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end

