/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */
package main

import (
	"fmt"
)

// @lc code=start
func jump(nums []int) int {
	/*
		贪心算法
		题目保证可达末尾, 因此不需要检查当前位置是否可达(i <= right_most)
	*/
	length := len(nums)
	end := 0        //存储本次跳跃可达的最远位置
	right_most := 0 //存储本次跳跃的区间内的所有位置可达的最远位置
	steps := 0
	for i := 0; i < length-1; i++ {
		if right_most < i+nums[i] {
			right_most = i + nums[i]
		}
		// 可以提前结束
		if right_most >= length-1 {
			steps++
			break
		}
		if i == end {
			// 遍历完了本次跳跃的区间所有元素, 右边界更新为最大可达位置
			// 相当于自动选择了最有的跳跃方案
			end = right_most
			steps++
		}
	}
	return steps
}

// @lc code=end

func main() {
	nums := []int{2, 3, 1, 2, 4, 2, 3}
	fmt.Println(jump(nums))
}
