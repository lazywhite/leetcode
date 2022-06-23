/*

每瓶汽水单价2元
规则
	4个瓶盖可以换1瓶汽水
	2个空瓶可以换1瓶汽水

30元可以喝到多少瓶汽水?
*/

package main

import "fmt"

func main() {
	fmt.Println(countBottle(30))
}

func countBottle(n int) int {
	unitPrice := 2

	lid := n / unitPrice
	empty := n / unitPrice
	total := n / unitPrice

	newBottle := 0 // 兑换到的新瓶子

	// 注意准入条件
	for lid+newBottle >= 4 || empty+newBottle >= 2 {
		// 1. drink
		lid += newBottle
		empty += newBottle

		// 2. 兑换
		newBottle = lid/4 + empty/2
		total += newBottle

		lid %= 4
		empty %= 2
	}

	return total
}
