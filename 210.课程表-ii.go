/*
 * @lc app=leetcode.cn id=210 lang=golang
 *
 * [210] 课程表 II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	// 构造有向图, 进行拓扑排序, 使用深度优先算法
	var (
		// 有向边(a, b), 如果b依赖a, 则 a -> b
		edges = make([][]int, numCourses) // 下标为课程编号, 存储其指向的课程
		// 顶点的入度
		indegree = make([]int, numCourses)
		result   = make([]int, 0)
	)

	// 构造有向边和入度
	for _, p := range prerequisites {
		edges[p[1]] = append(edges[p[1]], p[0])
		indegree[p[0]]++
	}
	q := []int{}
	// 将入度为0的节点放入队列
	for i, v := range indegree {
		if v == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		// 每次删除队列头部元素, 并将其指向的顶点的入度减一
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
	if len(result) != numCourses {
		return []int{}
	}
	return result
}

// @lc code=end

