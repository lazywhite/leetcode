/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	set1 := map[int]struct{}{}
	set2 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	for _, v := range nums2 {
		_, ok1 := set1[v]
		_, ok2 := set2[v]
		if ok1 && !ok2 {
			set2[v] = struct{}{}
		}
	}
	result := []int{}
	for k := range set2 {
		result = append(result, k)
	}
	return result
}

// @lc code=end

