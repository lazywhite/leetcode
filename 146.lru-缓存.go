/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存
 */

package main

import "fmt"

func main() {
	c := Constructor(2)
	c.Put(2, 1)
	c.Put(1, 1)
	c.Put(2, 3)
	c.Put(4, 1)
	c.Get(1)
	c.Get(2)
}

// @lc code=start
type Node struct {
	Key   int
	Value int
	Pre   *Node
	Next  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("Node{key=%d}", n.Key)
}

type LRUCache struct {
	Size     int
	Capacity int
	Head     *Node
	Tail     *Node
	Cache    map[int]*Node
}

/*
	1. hashtable缓存
	2. 双向链表
	3. 最新访问过的, 放在tail
	4. 超出capacity, 删掉head
*/
func Constructor(capacity int) LRUCache {
	head := NewNode(-1, 0)
	tail := NewNode(-100, 0)
	c := LRUCache{
		Size:     0,
		Capacity: capacity,
		Head:     head,
		Tail:     tail,
		Cache:    make(map[int]*Node, 0),
	}
	head.Next = tail
	tail.Pre = head
	return c
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

// 首先将其从原位置删除, 然后添加到tail
func (this *LRUCache) move2Tail(n *Node) {
	this.removeNode(n)
	this.addTail(n)
}

// 将node加到tail
func (this *LRUCache) addTail(n *Node) {
	pre := this.Tail.Pre
	pre.Next = n
	this.Tail.Pre = n
	n.Pre = pre
	n.Next = this.Tail
}

// 删除node
func (this *LRUCache) removeNode(n *Node) {
	n.Pre.Next = n.Next
	n.Next.Pre = n.Pre
}

// 如果不存在返回-1, 存在需要将其放在tail
func (this *LRUCache) Get(key int) int {
	n, ok := this.Cache[key]
	if !ok {
		return -1
	}
	this.move2Tail(n)
	return n.Value
}

func (this *LRUCache) Put(key int, value int) {
	// 如果已存在, 需要更新其value, 然后移动到tail
	if n, ok := this.Cache[key]; ok {
		n.Value = value
		this.move2Tail(n)
	} else {
		// 不存在就创建新node, 并添加到tail, 添加cache
		node := NewNode(key, value)
		this.Cache[key] = node
		this.addTail(node)
		if this.Size == this.Capacity {
			// 当前size等于capacity, 需要删除head
			// 1. 从cache删除
			delete(this.Cache, this.Head.Next.Key)
			// 2. 删除node
			this.removeNode(this.Head.Next)
		} else {
			this.Size++
		}
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
