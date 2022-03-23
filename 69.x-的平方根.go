/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {

	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		//中间数平方小于n, 左边界加1
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			// 要省略小数部分, 因此右边界不保存结果
			right = mid - 1
		}
	}
	return ans

}

// @lc code=end

