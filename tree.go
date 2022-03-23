//
// Copyright (C) 2022 lazywhite <lazywhite@qq.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
	}
	n1 := &TreeNode{
		Val: 2,
	}
	n2 := &TreeNode{
		Val: 3,
	}
	root.Right = n1
	n1.Left = n2

	fmt.Println(preorderTraversal(root))
}
