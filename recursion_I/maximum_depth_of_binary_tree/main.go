package main

// TreeNode implement tree data struct
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + maxDepth(root.Left)
	right := 1 + maxDepth(root.Right)
	if right > left {
		return right
	}
	return left
}

func main() {

}
