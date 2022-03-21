/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	bitmap := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		bitmap[nums[i]] = 1
	}

	result := make([]int, 0)
	for i := 1; i < len(bitmap); i++ {
		if bitmap[i] != 1 {
			result = append(result, i)
		}
	}
	return result

}

// @lc code=end

