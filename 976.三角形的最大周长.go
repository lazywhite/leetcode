/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长

 * 先从小到大排序, 然后从尾部遍历, 满足 a + b > c即可
 */

// @lc code=start
func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

// @lc code=end

