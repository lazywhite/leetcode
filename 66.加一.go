/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+plus > 9 {
			digits[i] = 0
			plus = 1
		} else {
			digits[i] = digits[i] + plus
			plus = 0
			break
		}
	}
	if plus > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits

}

// @lc code=end

