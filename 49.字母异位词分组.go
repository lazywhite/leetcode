/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 使用array(not slice)作为map key
	m := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, c := range str {
			cnt[c-'a']++
		}
		m[cnt] = append(m[cnt], str)
	}
	result := make([][]string, 0, len(m))
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// @lc code=end

