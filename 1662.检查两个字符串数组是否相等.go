/*
 * @lc app=leetcode.cn id=1662 lang=golang
 *
 * [1662] 检查两个字符串数组是否相等
 *
 * https://leetcode-cn.com/problems/check-if-two-string-arrays-are-equivalent/description/
 *
 * algorithms
 * Easy (80.96%)
 * Likes:    23
 * Dislikes: 0
 * Total Accepted:    24.4K
 * Total Submissions: 30.1K
 * Testcase Example:  '["ab", "c"]\n["a", "bc"]'
 *
 * 给你两个字符串数组 word1 和 word2 。如果两个数组表示的字符串相同，返回 true ；否则，返回 false 。
 *
 * 数组表示的字符串 是由数组中的所有元素 按顺序 连接形成的字符串。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：word1 = ["ab", "c"], word2 = ["a", "bc"]
 * 输出：true
 * 解释：
 * word1 表示的字符串为 "ab" + "c" -> "abc"
 * word2 表示的字符串为 "a" + "bc" -> "abc"
 * 两个字符串相同，返回 true
 *
 * 示例 2：
 *
 *
 * 输入：word1 = ["a", "cb"], word2 = ["ab", "c"]
 * 输出：false
 *
 *
 * 示例 3：
 *
 *
 * 输入：word1  = ["abc", "d", "defg"], word2 = ["abcddefg"]
 * 输出：true
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1
 * 1
 * 1
 * word1[i] 和 word2[i] 由小写字母组成
 *
 *
 */

// @lc code=start
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	s1 := new(strings.Builder)
	s2 := new(strings.Builder)
	for _, v := range word1 {
		s1.WriteString(v)
	}
	for _, v := range word2 {
		s2.WriteString(v)
	}

	str1 := s1.String()
	str2 := s2.String()

	if len(str1) != len(str2) {
		return false
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			return false
		}
	}
	return true

}

// @lc code=end

