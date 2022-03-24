/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	size := len(nums)
	bitmap := make([]int, size+1)

	for i := range nums {
		bitmap[nums[i]] = 1
	}

	for i := range bitmap {
		if bitmap[i] == 0 {
			return i
		}
	}
	return -1

}

// @lc code=end

