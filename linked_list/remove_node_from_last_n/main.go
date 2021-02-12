package main

import "fmt"

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func getLength(listNode *ListNode) int {
	head := listNode
	len := 0
	for head != nil {
		len++
		head = head.Next
	}
	return len
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := getLength(head)
	if len == 1 {
		return nil
	}
	if len-n <= 0 {
		return head.Next
	}
	target := len - n - 1
	l := head
	for i := 0; i < target; i++ {
		l = l.Next
	}
	fmt.Println()
	if n == 1 {
		l.Next = nil
		return head
	}

	l.Next = l.Next.Next
	return head
}

func PrintAll(listNode *ListNode) {
	head := listNode
	fmt.Print("[ ")
	for head != nil {
		fmt.Printf("%v", head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(", ")
		}

	}
	fmt.Printf(" ] \n")
}

func main() {
	nodeB5 := &ListNode{5, nil}
	nodeB4 := &ListNode{4, nodeB5}
	nodeB3 := &ListNode{3, nodeB4}
	nodeB2 := &ListNode{2, nodeB3}
	nodeB := &ListNode{1, nodeB2}
	// PrintAll(nodeB)
	// fmt.Println("---")
	PrintAll(removeNthFromEnd(nodeB, 2))
	fmt.Println("---")

	nodeA2 := &ListNode{2, nil}
	nodeA := &ListNode{1, nodeA2}
	PrintAll(removeNthFromEnd(nodeA, 2))
	fmt.Println("---")

}
