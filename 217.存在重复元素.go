/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	cache := map[int]int{}

	for i:=0;i<len(nums);i++{
		if _, ok:= cache[nums[i]];ok{
			return true
		}else{
			cache[nums[i]] = 1
		}
	}
	return false
}
// @lc code=end

