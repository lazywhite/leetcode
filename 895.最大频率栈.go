/*
 * @lc app=leetcode.cn id=895 lang=golang
 *
 * [895] 最大频率栈
 */

// @lc code=start
type FreqStack struct {
	MaxFreq  int
	Freq2Val map[int][]int // 存储频率与元素stack的映射
	Val2Freq map[int]int   // 存储元素与频率的映射
}

func Constructor() FreqStack {
	s := FreqStack{
		MaxFreq: 0,
	}
	s.Freq2Val = make(map[int][]int, 0)
	s.Val2Freq = make(map[int]int, 0)
	return s
}

func (this *FreqStack) Push(val int) {
	v := this.Val2Freq[val]

	freq := v + 1
	this.Val2Freq[val] = freq

	if this.MaxFreq < freq {
		this.MaxFreq = freq
	}

	if _, ok := this.Freq2Val[freq]; ok {
		this.Freq2Val[freq] = append(this.Freq2Val[freq], val)
	} else {
		this.Freq2Val[freq] = []int{val}
	}
}

func (this *FreqStack) Pop() int {
	if this.MaxFreq == 0 {
		return -1
	}
	// 查到最大频率
	max := this.MaxFreq

	// 查到此频率下对应的stack
	stack := this.Freq2Val[max]

	// 被pop的是最后一个, 它最晚被入栈
	elem := stack[len(stack)-1]

	// 移除此元素
	this.Freq2Val[max] = this.Freq2Val[max][:len(stack)-1]

	// 更新元素与频率映射
	this.Val2Freq[elem] -= 1

	// 更新最大频率
	if len(this.Freq2Val[max]) == 0 {
		// 如果存在max:stack映射, 则一定存在max-1:stack映射
		this.MaxFreq--
	}
	return elem

}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
// @lc code=end

