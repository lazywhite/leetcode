/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {

	min := math.MaxInt32
	max := 0 // default 0
	for i := 0; i < len(prices); i++ {
		cur := prices[i]
		if cur < min {
			min = cur
		}
		profit := cur - min
		if profit > max {
			max = profit
		}
	}
	return max
}

// @lc code=end

