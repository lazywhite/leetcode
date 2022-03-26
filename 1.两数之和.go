/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		subtract := target - current
		if v, ok := cache[subtract]; ok {
			return []int{v, i}
		}
		cache[current] = i
	}
	return []int{}
}

// @lc code=end
