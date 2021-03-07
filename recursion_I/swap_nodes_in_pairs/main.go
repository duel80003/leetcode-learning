package main

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	h := head.Next
	head.Next = swapPairs(head.Next.Next)
	h.Next = head
	return h
}

func printAll(listNode *ListNode) {
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
	node4 := &ListNode{4, nil}
	node3 := &ListNode{3, node4}
	node2 := &ListNode{2, node3}
	node1 := &ListNode{1, node2}
	printAll(node1)
	printAll(swapPairs(node1))
}
