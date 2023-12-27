/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	m := make(map[int]bool, len(nums))
	for _, i := range nums {
		m[i] = true
	}
	longest := 0
	for k := range m {
		if !m[k-1] { //上一个数字不存在, 表明当前数字是区间的起点
			cur := k + 1
			length := 1
			for m[cur] {
				length++
				cur++
			}
			if longest < length {
				longest = length
			}
		}
	}
	return longest
}

// @lc code=end

