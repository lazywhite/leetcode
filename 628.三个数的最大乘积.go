/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {

	min1, min2 := math.MaxInt64, math.MaxInt64
	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64

	for _, v := range nums {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
	}

	return max(min1*min2*max1, max1*max2*max3)

}

func max (x, y int)  int {
	if x > y {

	return x
	}
	return y
}

// @lc code=end

