package main

import "fmt"

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseLinkedList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var h *ListNode
	for head.Next != nil {
		tmp := head.Next
		head.Next = tmp.Next

		if h == nil {
			tmp.Next = head
			h = tmp
		} else {
			tmp.Next = h
			h = tmp
		}

	}
	return h
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
	nodeB5 := &ListNode{5, nil}
	nodeB4 := &ListNode{4, nodeB5}
	nodeB3 := &ListNode{3, nodeB4}
	nodeB2 := &ListNode{2, nodeB3}
	r := reverseLinkedList(nodeB2)
	PrintAll(r)

}
