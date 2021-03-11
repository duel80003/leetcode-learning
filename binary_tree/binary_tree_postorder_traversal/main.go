package main

import "fmt"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	r := []int{}
	if root == nil {
		return r
	}
	stack := []*TreeNode{}

	stack = append(stack, root)

	for len(stack) != 0 {
		root = stack[len(stack)-1]
		if root.Left == nil && root.Right == nil {
			stack = stack[:len(stack)-1]
			r = append(r, root.Val)
		}
		if root.Right != nil {
			stack = append(stack, root.Right)
			root.Right = nil
		}

		if root.Left != nil {
			stack = append(stack, root.Left)
			root.Left = nil
		}
	}
	return r
}

func main() {
	t1 := &TreeNode{1, nil, nil}
	t2 := &TreeNode{2, nil, nil}
	t3 := &TreeNode{3, nil, nil}
	t1.Right = t2
	t2.Left = t3
	fmt.Println(postorderTraversal(t1))
}
