/*
 * @lc app=leetcode.cn id=797 lang=golang
 *
 * [797] 所有可能的路径
 */
package main

import "fmt"

// @lc code=start
func allPathsSourceTarget(graph [][]int) [][]int {
	dest := len(graph) - 1

	result := [][]int{}
	var dfs func(int, []int)
	dfs = func(node int, past []int) {
		past = append(past[:], node)
		if node == dest {
			result = append(result, past)
		}
		for _, n := range graph[node] {
			dfs(n, past)
		}
	}
	//TODO: 结果异常
	dfs(0, []int{})
	return result
}

// @lc code=end

func main() {
	graph := [][]int{
		{3, 1},
		{4, 6, 7, 2, 5},
		{4, 6, 3},
		{6, 4},
		{7, 6, 5},
		{6},
		{7},
		{},
	}
	fmt.Println(allPathsSourceTarget(graph))
}
