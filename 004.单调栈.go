//
// Copyright (C) 2022 lazywhite <lazywhite@qq.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
)

func main() {
	nums := []int{10, 3, 7, 4, 12}
	ans := monoStack(nums)
	fmt.Println(nums)
	fmt.Println(ans)
}

/*
1. 下一个, 从前向后遍历
   上一个, 从后向前遍历
2. smarller 弹出比v大的
2. larger 弹出比v小的

*/

func monoStack(nums []int) []int {
	// stack存储num的index
	stack := make([]int, 0)
	ans := make([]int, len(nums))

	for i, v := range nums {
		if len(stack) == 0 {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && nums[stack[len(stack)-1]] > v {
				// 更新结果
				peek := stack[len(stack)-1]
				ans[peek] = v
				stack = stack[:len(stack)-1]
			}
			// 入栈
			stack = append(stack, i)
		}
	}

	// stack不为空
	if len(stack) > 0 {
		for _, idx := range stack {
			ans[idx] = -1
		}
	}
	return ans
}
