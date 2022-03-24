/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串

	s可以映射为t
	t也可以映射为s
	映射关系可以不同
*/

// @lc code=start
func isIsomorphic(s string, t string) bool {
	size1 := len(s)
	size2 := len(t)
	if size1 != size2 {
		return false
	}
	if s == "" && t == "" {
		return true
	}
	s2t := make(map[byte]byte, 0)
	t2s := make(map[byte]byte, 0)

	for i := range s {
		schar := s[i]
		tchar := t[i]

		if v, ok := s2t[schar]; ok {
			if tchar != v {
				return false
			}
		} else {
			s2t[schar] = tchar
		}
		if v, ok := t2s[tchar]; ok {
			if schar != v {
				return false
			}
		} else {
			t2s[tchar] = schar
		}
	}
	return true
}

// @lc code=end

