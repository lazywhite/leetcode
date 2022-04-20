/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode-cn.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (71.22%)
 * Likes:    175
 * Dislikes: 0
 * Total Accepted:    118.9K
 * Total Submissions: 166.7K
 * Testcase Example:  '3'
 *
 * 给你一个整数 n ，找出从 1 到 n 各个整数的 Fizz Buzz 表示，并用字符串数组 answer（下标从 1
 * 开始）返回结果，其中：
 *
 *
 * answer[i] == "FizzBuzz" 如果 i 同时是 3 和 5 的倍数。
 * answer[i] == "Fizz" 如果 i 是 3 的倍数。
 * answer[i] == "Buzz" 如果 i 是 5 的倍数。
 * answer[i] == i （以字符串形式）如果上述条件全不满足。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["1","2","Fizz"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 5
 * 输出：["1","2","Fizz","4","Buzz"]
 *
 *
 * 示例 3：
 *
 *
 * 输入：n = 15
 *
 * 输出：["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 10^4
 *
 *
 */

// @lc code=start
func fizzBuzz(n int) []string {

	ans := make([]string, n)

	// 标记3的倍数
	i := 1
	for i <= n {
		multi := i * 3
		if multi > n {
			break
		}
		idx := multi - 1

		if ans[idx] == "" {
			ans[idx] = "Fizz"
		}
		i++
	}

	// 标记5的倍数
	i = 1
	for i <= n {
		multi := i * 5
		if multi > n {
			break
		}
		idx := multi - 1

		if ans[idx] == "" {
			ans[idx] = "Buzz"
		} else if ans[idx] == "Fizz" {
			ans[idx] = "FizzBuzz"
		}

		i++
	}

	// 空字符串替换为idx+1
	for j := range ans {
		if ans[j] == "" {
			ans[j] = strconv.Itoa(j + 1)
		}
	}
	return ans
}

// @lc code=end

