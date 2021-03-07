package main

import "fmt"

// TreeNode implement tree data struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	root.Left = searchBST(root.Left, val)
	if root.Left != nil {
		return root.Left
	}
	root.Right = searchBST(root.Right, val)
	return root.Right
}

func printTree(t *TreeNode) {
	fmt.Printf("[ %d, %d, %d ] \n", t.Val, t.Left.Val, t.Right.Val)
}

func main() {
	left := &TreeNode{
		2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, nil, nil},
	}
	right := &TreeNode{7, nil, nil}
	r := &TreeNode{1, left, right}
	printTree(searchBST(r, 2))
}
