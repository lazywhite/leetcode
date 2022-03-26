/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if left != i {
				swap(nums, i, left)
			}
			left++
		}
	}
}
func swap(nums []int, x, y int) {
	nums[x], nums[y] = nums[y], nums[x]
}

// @lc code=end

