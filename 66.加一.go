/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	plus := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + plus
		digits[i] = sum % 10
		plus = sum / 10
		if plus != 1 {
			break
		}
	}

	// 最高位有进位
	if plus == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits

}

// @lc code=end

