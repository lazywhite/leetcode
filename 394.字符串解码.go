/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "3[a2[c]]"
	fmt.Println(decodeString(x))
}

// @lc code=start
func decodeString(s string) string {
	stack := []string{}
	i := 0
	for i < len(s) {
		char := s[i]
		if char == '[' {
			stack = append(stack, string(char))
			i++
			continue
		}
		if char >= '0' && char <= '9' {
			number := getNumber(s, &i)
			stack = append(stack, number)
			continue
		}
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			str := getString(s, &i)
			stack = append(stack, str)
			continue
		}
		if char == ']' {
			tempStr := ""
			// pop str until "["
			for len(stack) > 0 {
				c := stack[len(stack)-1]
				if c != "[" {
					tempStr = c + tempStr
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}

			stack = stack[:len(stack)-1] // pop "["
			num := stack[len(stack)-1]   // get number
			stack = stack[:len(stack)-1] // pop number
			rstr := repeateStr(tempStr, num)
			// push back to stack
			stack = append(stack, rstr)
			i++
			continue
		}
	}
	final := ""
	for i := range stack {
		final += stack[i]
	}
	return final
}

func repeateStr(str string, times string) string {
	result := ""
	t, _ := strconv.Atoi(times)
	for i := 0; i < t; i++ {
		result += str
	}
	return result
}

func getNumber(s string, start *int) string {
	x := ""
	for *start < len(s) {
		char := s[*start]
		if char >= '0' && char <= '9' {
			x += string(char)
			*start++
		} else {
			break
		}
	}
	return x
}

func getString(s string, start *int) string {
	x := ""
	for *start < len(s) {
		char := s[*start]
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			x += string(char)
			*start++
		} else {
			break
		}
	}
	return x
}

// @lc code=end
