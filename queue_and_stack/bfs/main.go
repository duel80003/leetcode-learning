package main

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bfs(root *TreeNode) []int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	result := []int{}
	return BFSUtil(queue, result)
}

// BFSUtil implement bfs
func BFSUtil(queue []*TreeNode, res []int) []int {
	if len(queue) == 0 {
		return res
	}
	res = append(res, queue[0].Val)
	if queue[0].Right != nil {
		queue = append(queue, queue[0].Right)
	}
	if queue[0].Left != nil {
		queue = append(queue, queue[0].Left)
	}
	return BFSUtil(queue, res)
}

func main() {

}
