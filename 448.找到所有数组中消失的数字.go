/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 */

// @lc code=start
func findDisappearedNumbers(nums []int) (ans []int) {
	/*
		下标范围(0, n-1)
		值范围(1, n)
	*/
	n := len(nums)
	for _, v := range nums {
		v = v % n // 取到增加前的值
		idx = v - 1 // 此值对应的index
		nums[idx] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return

}

// @lc code=end

