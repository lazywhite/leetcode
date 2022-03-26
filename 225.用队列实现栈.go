/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

/*

 创建两个队列
push
	1. 往2号队列push
	2. 将1号队列的全部pop, push到2号queue
 	3. 然后互换1, 2号队列
pop
	直接从1号队列pop
*/

// @lc code=start
type MyStack struct {
}

func Constructor() MyStack {

}

func (this *MyStack) Push(x int) {

}

func (this *MyStack) Pop() int {

}

func (this *MyStack) Top() int {

}

func (this *MyStack) Empty() bool {

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end

