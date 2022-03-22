/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	// nums并不包含0, 创建包含n+1个元素的slice
	bitmap := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		bitmap[nums[i]] += 1
	}

	result := make([]int, 0)
	// 跳过index=0
	for i := 1; i < len(bitmap); i++ {
		if bitmap[i] == 0 {
			result = append(result, i)
		}
	}
	return result

}

// @lc code=end

