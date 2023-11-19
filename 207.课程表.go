/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		edges    = make([][]int, numCourses)
		indegree = make([]int, numCourses)
		result   = make([]int, 0)
	)

	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indegree[p[0]]++
	}
	q := []int{}
	for i, v := range indegree {
		if v == 0 {
			q = append(q, i)
		}
	}
	/*
		如果存在环, 则这个环上所有顶点的入度都大于0, 无法放进q, 从而退出循环, 结果的数量肯定小于numCourses
	*/
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		result = append(result, c)
		for _, v := range edges[c] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return len(result) == numCourses
}

// @lc code=end

