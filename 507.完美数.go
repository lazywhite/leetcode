/*
 * @lc app=leetcode.cn id=507 lang=golang
 *
 * [507] 完美数
 */

// @lc code=start
func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	sum := 1
	for d := 2; d*d <= num; d++ {
		if num%d == 0 {
			sum += d
			if d*d < num {
				sum += num / d
			}
		}
	}
	return sum == num

}

// @lc code=end

