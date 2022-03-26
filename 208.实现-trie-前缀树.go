/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	Head *Node
}

type Node struct {
	Char     byte
	IsEnd    bool
	Children map[byte]*Node
}

func Constructor() Trie {
	head := &Node{
		Char:  0,
		Children: make(map[byte]*Node, 0),
		IsEnd: false,
	}
	t := Trie{
		Head:     head,
	}
	return t
}

func (this *Trie) Insert(word string) {
	curNode := this.Head
	for i := range word {
		char := word[i]
		if node, ok := curNode.Children[char]; !ok {
			n := &Node{
				Char: char,
				Children: make(map[byte]*Node, 0),
			}
			curNode.Children[char] = n
			curNode = n

		} else {
			curNode = node
		}
	}

	curNode.IsEnd = true
}

func (this *Trie) Search(word string) bool {
	curNode := this.Head
	for i := range word {
		char := word[i]
		if node, ok := curNode.Children[char]; !ok {
			return false
		} else {
			curNode = node
		}
	}
	return curNode.IsEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	curNode := this.Head
	for i := range prefix {
		char := prefix[i]
		if node, ok := curNode.Children[char]; !ok {
			return false
		} else {
			curNode = node
		}
	}
	return true

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

