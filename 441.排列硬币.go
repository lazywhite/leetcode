/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] 排列硬币
 */

/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println(arrangeCoins(1804289383))
}
*/

// @lc code=start
func arrangeCoins(n int) int {
	if n == 1 {
		return 1
	}
	// 仅需要记忆上一个数, 因此可以省去数组
	preSum := 0
	i := 1
	for i <= n {
		curSum := i + preSum
		if curSum == n {
			return i
		}
		if curSum > n {
			return i - 1
		}
		preSum = curSum
		i++
	}
	return 0
}

// @lc code=end

/*
func bak(n int) int {

	// 每行的个数, 递增数列, 退化成index
	line := make([]int, n)
	// 每行的总数
	sum := make([]int, n)
	line[1] = 1
	sum[1] = 1

	k := 2

	for k < n {
		line[k] = line[k-1] + 1
		sum[k] = sum[k-1] + line[k]
		if sum[k] < n {
			k++
		}
		if sum[k] == n {
			break
		}
		if sum[k] > n {
			k--
			break
		}
	}
	return k
}
*/
