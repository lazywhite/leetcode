/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {

	lastSum := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		curSum := nums[i]
		if lastSum > 0 {
			curSum = curSum + lastSum
		}
		if maxSum < curSum {
			maxSum = curSum
		}
		lastSum = curSum
	}
	return maxSum

}

// @lc code=end

