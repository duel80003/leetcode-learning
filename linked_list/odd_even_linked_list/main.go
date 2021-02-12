package main

import "fmt"

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	index := 1
	odd := &ListNode{-1, nil}
	even := &ListNode{-1, nil}
	oddT := odd
	evenT := even
	for head != nil {
		if index%2 == 0 {
			evenT.Next = head
			evenT = evenT.Next
		} else {
			oddT.Next = head
			oddT = oddT.Next
		}
		head = head.Next
		index++
	}
	evenT.Next = nil
	oddT.Next = even.Next
	return odd.Next
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
	nodeB7 := &ListNode{7, nil}
	nodeB6 := &ListNode{4, nodeB7}
	nodeB5 := &ListNode{6, nodeB6}
	nodeB4 := &ListNode{5, nodeB5}
	nodeB3 := &ListNode{3, nodeB4}
	nodeB2 := &ListNode{1, nodeB3}
	nodeB := &ListNode{2, nodeB2}
	r := oddEvenList(nodeB)
	PrintAll(r)

	nodeC6 := &ListNode{4, nil}
	nodeC5 := &ListNode{6, nodeC6}
	nodeC4 := &ListNode{5, nodeC5}
	nodeC3 := &ListNode{3, nodeC4}
	nodeC2 := &ListNode{1, nodeC3}
	nodeC := &ListNode{2, nodeC2}

	rc := oddEvenList2(nodeC)
	PrintAll(rc)
}
