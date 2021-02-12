package main

import "fmt"

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	h := head
	for h != nil && h.Next != nil {
		if h.Next.Val == val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}
	if head.Val == val {
		return head.Next
	}
	return head
}

// PrintAll listnode element
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
	nodeB8 := &ListNode{5, nil}
	// nodeB7 := &ListNode{5, nodeB8}
	// nodeB6 := &ListNode{1, nodeB7}
	// nodeB5 := &ListNode{5, nodeB6}
	// nodeB4 := &ListNode{4, nodeB5}
	// nodeB3 := &ListNode{3, nodeB4}
	// nodeB2 := &ListNode{2, nodeB3}
	nodeB := &ListNode{5, nodeB8}
	r := removeElements(nodeB, 5)
	fmt.Println("\n---")
	PrintAll(r)
}
