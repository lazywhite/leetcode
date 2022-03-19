/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		middle := (low + high) / 2
		val := nums[middle]
		if target == val {
			return middle
		}
		if target < val {
			high = middle - 1
		} else {
			low = middle + 1
		}

	}
	return -1
}

// @lc code=end

