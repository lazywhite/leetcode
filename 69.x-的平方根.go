/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {

	left, right := 0, x
	ans := -1 // 永远存储mid的值
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		}
		// 只有左边界发生改变, 才存储mid值
		if mid*mid < x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans

}

// @lc code=end

