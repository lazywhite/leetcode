/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {

	// 绝不会有重复计算, 因此无需缓存
	// cache := make(map[int]float64, 0)

	if n < 0 {
		return 1.0 / quick(x, -n)
	}
	if n == 0 {
		return 1
	}

	return quick(x, n)
}
func quick(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := quick(x, n/2)

	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

// @lc code=end
