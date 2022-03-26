/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

/*
	栈顶的元素未出栈时, 它下方的元素是固定的, 所以最小值也是固定的
	因此构造一个辅助栈, 用来存储当前的最小值, 跟栈的元素是一一对应的
*/
type MinStack struct {
	Stack []int
	Min   []int
}

func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}

}

func (this *MinStack) Push(val int) {
	if len(this.Min) == 0 {
		this.Min = append(this.Min, val)
	} else {
		this.Min = append(this.Min, min(this.Min[len(this.Min)-1], val))
	}
	this.Stack = append(this.Stack, val)

}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.Min = this.Min[:len(this.Min)-1]

}

func (this *MinStack) Top() int {

	if len(this.Stack) > 0 {
		return this.Stack[len(this.Stack)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.Min) > 0 {
		return this.Min[len(this.Min)-1]
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

