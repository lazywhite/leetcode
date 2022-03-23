//
// Copyright (C) 2022 lazywhite <lazywhite@qq.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	root := &ListNode{
		Val: 1,
	}
	n1 := &ListNode{
		Val: 2,
	}
	n2 := &ListNode{
		Val: 3,
	}
	root.Next = n1
	n1.Next = n2

	fmt.Println(reverseList(root))
}
