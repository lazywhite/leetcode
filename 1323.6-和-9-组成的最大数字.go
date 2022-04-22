/*
 * @lc app=leetcode.cn id=1323 lang=golang
 *
 * [1323] 6 和 9 组成的最大数字
 *
 * https://leetcode-cn.com/problems/maximum-69-number/description/
 *
 * algorithms
 * Easy (75.25%)
 * Likes:    59
 * Dislikes: 0
 * Total Accepted:    28.4K
 * Total Submissions: 37.7K
 * Testcase Example:  '9669'
 *
 * 给你一个仅由数字 6 和 9 组成的正整数 num。
 *
 * 你最多只能翻转一位数字，将 6 变成 9，或者把 9 变成 6 。
 *
 * 请返回你可以得到的最大数字。
 *
 *
 *
 * 示例 1：
 *
 * 输入：num = 9669
 * 输出：9969
 * 解释：
 * 改变第一位数字可以得到 6669 。
 * 改变第二位数字可以得到 9969 。
 * 改变第三位数字可以得到 9699 。
 * 改变第四位数字可以得到 9666 。
 * 其中最大的数字是 9969 。
 *
 *
 * 示例 2：
 *
 * 输入：num = 9996
 * 输出：9999
 * 解释：将最后一位从 6 变到 9，其结果 9999 是最大的数。
 *
 * 示例 3：
 *
 * 输入：num = 9999
 * 输出：9999
 * 解释：无需改变就已经是最大的数字了。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= num <= 10^4
 * num 每一位上的数字都是 6 或者 9 。
 *
 *
 */

// 使用字符串
func bak(num int) int {
	s := strconv.Itoa(num)
	b := []byte(s)
	for i := range b {
		if b[i] == '6' {
			b[i] = '9'
			break
		}
	}

	ans, _ := strconv.Atoi(string(b))
	return ans
}

// @lc code=start
func maximum69Number(num int) int {
	i := 1
	match := 0

	last := 0

	for s := num; s != 0; {
		last = s % 10
		if last == 6 {
			match = 3 * i
		}
		s /= 10
		i *= 10
	}
	return match + num
}

// @lc code=end

