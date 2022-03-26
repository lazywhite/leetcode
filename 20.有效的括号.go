/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if s == "" {
		return false
	}
	matches := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	stack := []string{}

	for i := range s {
		char := string(s[i])
		if v, ok := matches[char]; ok {
			stack = append(stack, v)
		} else {
			if len(stack) == 0 {
				return false
			} else {
				if stack[len(stack)-1] != char {
					return false
				} else {
					stack = stack[:len(stack)-1]
				}
			}
		}
	}
	return len(stack) == 0
}

// @lc code=end

