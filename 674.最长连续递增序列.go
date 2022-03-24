/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {

	currentSize := 1
	maxSize := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			currentSize = 1
		} else {
			currentSize++
			maxSize = max(maxSize, currentSize)
		}
	}
	return maxSize
}
func max(x, y int) int{
	if x > y {
		return x
	}
		return y
}

// @lc code=end

