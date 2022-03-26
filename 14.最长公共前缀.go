/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	if len(strs) == 0 {
		return ""
	}
	return lcp(strs, 0, len(strs)-1)

}

func lcp(strs []string, left, right int) string {
	// left can never be bigger than right
	if left == right {
		return strs[left]
	}
	mid := (left + right) / 2
	leftPrefix := lcp(strs, left, mid)
	rightPrefix := lcp(strs, mid+1, right)
	return findCommonPrefix(leftPrefix, rightPrefix)
}

func findCommonPrefix(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	min := min(len1, len2)
	if min == 0 {
		return ""
	}

	start := 0
	for start <= min-1 {
		if str1[start] != str2[start] {
			break
		}
		start++
	}
	return str1[:start]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}

// @lc code=end

