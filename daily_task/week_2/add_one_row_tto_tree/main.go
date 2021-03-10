package main

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if d == 1 {
		return &TreeNode{
			Val:   v,
			Left:  root,
			Right: nil,
		}
	}
	dfs(root, 1, v, d)
	return root
}

func dfs(node *TreeNode, c, v, d int) {
	if node == nil {
		return
	}
	if c == d-1 {
		node.Left = &TreeNode{
			Val:   v,
			Left:  node.Left,
			Right: nil,
		}
		node.Right = &TreeNode{
			Val:   v,
			Left:  nil,
			Right: node.Right,
		}
		return
	}
	dfs(node.Left, c+1, v, d)
	dfs(node.Right, c+1, v, d)
}

func main() {

}
