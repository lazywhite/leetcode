/*
 * @lc app=leetcode.cn id=345 lang=golang
 *
 * [345] 反转字符串中的元音字母
 */

// @lc code=start
func reverseVowels(s string) string {
	c := map[string]int{
		"a": 1,
		"e": 1,
		"i": 1,
		"o": 1,
		"u": 1,
		"A": 1,
		"E": 1,
		"I": 1,
		"O": 1,
		"U": 1,
	}
	size := len(s)
	left := 0
	right := size - 1

	t := []byte(s)

	for left < right {
		for left < size {
			_, ok := c[string(t[left])]
			if ok {
				break
			}
			left++
		}
		for right >= 0 {
			_, ok := c[string(t[right])]
			if ok {
				break
			}
			right--
		}
		if left < right {
			swap(t, left, right)
			left++
			right--
		}
	}
	return string(t)
}

func swap(s []byte, x, y int) {
	s[x], s[y] = s[y], s[x]
}

// @lc code=end

