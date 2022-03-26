/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

/*
1. 任何数和 00 做异或运算，结果仍然是原来的数，即 a \oplus 0=aa⊕0=a。
2. 任何数和其自身做异或运算，结果是 00，即 a \oplus a=0a⊕a=0。
3. 异或运算满足交换律和结合律，即 a \oplus b \oplus a=b \oplus a \oplus a=b \oplus (a \oplus a)=b \oplus0=ba⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b
*/

// @lc code=start
func singleNumber(nums []int) int {
	var result int
	for i := range nums {
		result ^= nums[i]
	}
	return result

}

// @lc code=end

