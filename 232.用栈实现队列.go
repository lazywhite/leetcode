/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

/*
push
 	1. 先将1号stack的元素全部pop, 然后push到2号stack
 	2. 将元素push到1号stack
 	3. 将元素从2号stack全部pop, 然后push到1号stack

pop
	直接从1号stack pop
*/

// @lc code=start
type MyQueue struct {
}

func Constructor() MyQueue {

}

func (this *MyQueue) Push(x int) {

}

func (this *MyQueue) Pop() int {

}

func (this *MyQueue) Peek() int {

}

func (this *MyQueue) Empty() bool {

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

