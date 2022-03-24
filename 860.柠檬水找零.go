/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	stack5 := 0
	stack10 := 0
	// 无需存储面额20的信息, 只能增加, 无法减少
	//stack20 := 0

	for _, income := range bills {

		switch income {
		case 5:
			stack5++
		case 10:
			stack5--
			stack10++
		case 20:
			//stack20++
			if stack10 > 0 {
				stack5--
				stack10--
			} else {
				stack5 -= 3
			}
		}

		if stack5 < 0 {
			return false
		}

	}
	return true

}

// @lc code=end

