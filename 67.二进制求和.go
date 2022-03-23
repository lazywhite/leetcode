package main

func main() {

	a := "11"
	b := "1"
	addBinary(a, b)
}

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */


// 与[415], [66]相同

// @lc code=start
func addBinary(a string, b string) string {

	result := []byte{}
	// '0' = 48

	s1 := a // 长度最长的
	s2 := b
	if len(a) < len(b) {
		s1 = b
		s2 = a
	}

	// 进位
	carry := 0

	for i := 0; i < len(s1); i++ {
		idx1 := len(s1) - 1 - i
		idx2 := len(s2) - 1 - i
		i1 := int(s1[idx1]) - 48
		i2 := 0
		if idx2 >= 0 {
			i2 = int(s2[idx2]) - 48
		}

		left := (i1 + i2 + carry) % 2
		carry = (i1 + i2 + carry) / 2

		result = append([]byte{
			byte(left + 48),
		}, result...)
	}
	if carry == 1 {
		result = append([]byte{
			byte(1 + 48),
		}, result...)
	}
	return string(result)

}

// @lc code=end
