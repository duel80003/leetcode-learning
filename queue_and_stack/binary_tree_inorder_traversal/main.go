package main

// TreeNode for a binary tree node.
type TreeNode struct {
	Val   int
	Left  TreeNode
	Right TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	stack := []*TreeNode{}
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			cur = cur.Right
			for cur != nil {
				stack = append(stack, cur)
				cur = cur.Left
			}
		}
	}

	return result
}

func main() {

}
