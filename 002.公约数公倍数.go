//
// Copyright (C) 2022 lazywhite <lazywhite@qq.com>
//
// Distributed under terms of the MIT license.
//

package main

import "fmt"

func main() {

	//fmt.Println(getCommonDivisor(28, 14))
	fmt.Println(getCommonMultiple(25, 10))
}

// 最大公约数
// 更相减损术
func getCommonDivisor(a, b int) int {
	big := a
	small := b
	if a < b {
		big = b
		small = a
	}

	if big%small == 0 {
		return small
	}
	return getCommonDivisor(big-small, small)
}

// 最大公倍数
// 两数的乘积除以最大公约数
func getCommonMultiple(a, b int) int{

	multi := a * b
	divisor := getCommonDivisor(a, b)

    // 一定能整除
	return multi/divisor
}
