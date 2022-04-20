/*
 * @lc app=leetcode.cn id=925 lang=golang
 *
 * [925] 长按键入
 *
 * https://leetcode-cn.com/problems/long-pressed-name/description/
 *
 * algorithms
 * Easy (38.38%)
 * Likes:    228
 * Dislikes: 0
 * Total Accepted:    54K
 * Total Submissions: 140.6K
 * Testcase Example:  '"alex"\n"aaleex"'
 *
 * 你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
 *
 * 你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：name = "alex", typed = "aaleex"
 * 输出：true
 * 解释：'alex' 中的 'a' 和 'e' 被长按。
 *
 *
 * 示例 2：
 *
 *
 * 输入：name = "saeed", typed = "ssaaedd"
 * 输出：false
 * 解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= name.length, typed.length <= 1000
 * name 和 typed 的字符都是小写字母
 *
 *
 */
 package main
 func main(){
	 n := "vex"
	 r := "vext"
 }

// @lc code=start
func isLongPressedName(name string, typed string) bool {

	p1 := 0
	p2 := 0

	s1 := len(name)
	s2 := len(typed)

	// 遍历typed, 而不是name
	for p2 < s2 {
		if p1 < s1 && name[p1] == typed[p2] {
			p1++
			p2++
		} else {
			if p2 == 0 {
				return false
			} else {
				if typed[p2] == typed[p2-1] {
					p2++
				} else {
					return false
				}
			}
		}
	}
	return p1 == s1
}

// @lc code=end

