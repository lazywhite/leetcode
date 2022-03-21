/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

package main

import (
	"fmt"
)

func main() {
	x := []int{2, 3, 2}
	fmt.Println(rob(x))
}

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	/*
	由于是环形的, 所以3个房屋其实是挨着的, 所以取最大值即可
	*/
	if len(nums) == 3 {
		return max(max(nums[0], nums[1]), nums[2])
	}
	// 1. 偷了第一家, 不能偷第二家及最后一家
	num1 := nums[0] + robIt(nums, 2, len(nums)-2)
	// 2. 没偷第一家, 可以偷第二家到最后一家
	num2 := robIt(nums, 1, len(nums)-1)
	return max(num1, num2)
}

func robIt(nums []int, start, end int) int {
	if start == end {
		return nums[start]
	}
	if end-start == 1 {
		return max(nums[start], nums[end])
	}
	dp := make([]int, end+1)
	dp[start] = nums[start]
	dp[start+1] = max(nums[start], nums[start+1])

	for i := start + 2; i <= end; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[end]
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
