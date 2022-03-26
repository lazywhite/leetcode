/*
 * @lc app=leetcode.cn id=460 lang=golang
 *
 * [460] LFU 缓存
 */

// @lc code=start
type LFUCache struct {
	Size      int
	Capacity  int
	MinFreq   int
	Freq2Node map[int][]*Node // queue
	Key2Node  map[int]*Node   // 元素的访问次数

}
type Node struct {
	Key   int
	Value int
	Count int
}

func Constructor(capacity int) LFUCache {

}

func (this *LFUCache) Get(key int) int {

}

func (this *LFUCache) Put(key int, value int) {

}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

