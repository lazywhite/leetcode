/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 *
 * https://leetcode-cn.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (60.23%)
 * Likes:    167
 * Dislikes: 0
 * Total Accepted:    66.9K
 * Total Submissions: 111.1K
 * Testcase Example:  '"ab-cd"'
 *
 * 给你一个字符串 s ，根据下述规则反转字符串：
 *
 *
 * 所有非英文字母保留在原有位置。
 * 所有英文字母（小写或大写）位置反转。
 *
 *
 * 返回反转后的 s 。
 *
 *
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "ab-cd"
 * 输出："dc-ba"
 *
 *
 *
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "a-bC-dEf-ghIj"
 * 输出："j-Ih-gfE-dCba"
 *
 *
 *
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "Test1ng-Leet=code-Q!"
 * 输出："Qedo1ct-eeLg=ntse-T!"
 *
 *
 *
 *
 * 提示
 *
 *
 * 1 <= s.length <= 100
 * s 仅由 ASCII 值在范围 [33, 122] 的字符组成
 * s 不含 '\"' 或 '\\'
 *
 *
 */

// @lc code=start
func reverseOnlyLetters(s string) string {
	ans := []byte(s)
	size := len(ans)

	left := 0
	right := size - 1

	for left < right {
		for left < size && !isLetter(ans[left]) {
			left++
		}
		for right >=0 && !isLetter(ans[right]){
			right--
		}
		if left < right {
			swap(ans, left, right)
			left++
			right--
		}
	}

	return string(ans)
}
func swap(arr []byte, x, y int) {
	arr[x], arr[y] = arr[y], arr[x]
}

func isLetter(c byte) bool {
	if c >= byte('a') && c <= byte('z') {
		return true
	} else if c >= byte('A') && c <= byte('Z') {
		return true
	}
	return false
}

// @lc code=end

