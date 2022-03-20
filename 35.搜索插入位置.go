/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {

	return findIdx(nums, 0, len(nums)-1, target)

}

func findIdx(nums []int, left, right, target int) int {
	if left > right {
		return 0
	}
	if left == right {
		if target > nums[left] {
			return left + 1
		}
		return left
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if target < nums[mid] {
		// 要包含mid
		return findIdx(nums, left, mid, target)
	}
	// 不能包含mid
	return findIdx(nums, mid+1, right, target)
}

// @lc code=end

