/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	left, right := 0, 1

	for right < len(nums) {
		if nums[left] != nums[right] {
			if left != right {
				nums[left+1] = nums[right]
				left++
			} else {
				left++
			}
		}

		right++

	}
	// need size not index
	return left + 1

}

// @lc code=end

