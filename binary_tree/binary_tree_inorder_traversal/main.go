package main

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  TreeNode
	Right TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, node.Val)
			cur = node.Right
		}
	}
	return result
}

func inorderTraversalWithRecursive(root *TreeNode) []int {
	return recursive(root)
}

func recursive(node *TreeNode) {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	result = append(recursive(root.Left), root.Val)
	result = append(result, recursive(root.Right)...)
	return result
}

func main() {

}
