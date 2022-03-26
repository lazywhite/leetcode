/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长

 * 先从小到大排序, 然后从尾部遍历, 满足 a + b > c即可
 */

// @lc code=start
func largestPerimeter(nums []int) int {
	sort(nums)

	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func sort(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}

		}
	}
}

// @lc code=end

