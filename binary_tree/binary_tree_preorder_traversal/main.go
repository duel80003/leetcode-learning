package main

import "fmt"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// if root == nil {
	// 	return nil
	// }
	// head := root
	// r := []int{}
	// stack := []TreeNode{}
	// stack = append(stack, *root)
	// for stack != nil {
	// 	node := stack[:len(stack)-1]
	// 	if head.Left != nil {
	// 		r = append(r, head.Left.Val)
	// 	} else {

	// 	}
	// }
	// return r
	return nil
}

func main() {
	t1 := &TreeNode{1, nil, nil}
	t2 := &TreeNode{2, nil, nil}
	t3 := &TreeNode{3, nil, nil}
	t1.Right = t2
	t2.Left = t3
	fmt.Println(preorderTraversal(t1))
}
