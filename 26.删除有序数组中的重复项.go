/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}
	left := 1
	for right := 1; right < size; right++ {
		if nums[right] != nums[right-1] {
			// 直接用right值覆盖left
			nums[left] = nums[right]
			left++
		}
	}
	return left

}

// @lc code=end

