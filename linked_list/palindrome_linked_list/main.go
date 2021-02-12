package main

import (
	"fmt"
)

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func fillArray(head *ListNode) *[]int {
	h := head
	r := []int{}
	for h != nil {
		r = append(r, h.Val)
		h = h.Next
	}
	return &r
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	array := *fillArray(head)
	lenght := len(array)
	for i := lenght - 1; i >= 0; i-- {
		if array[i] != head.Val {
			return false
		}
		head = head.Next
	}
	return true
}

func reverseLinkedList(list *ListNode) *ListNode {
	if list == nil || list.Next == nil {
		return list
	}
	var h *ListNode
	head := list
	for head != nil && head.Next != nil {
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

func getPalindromeStart(list *ListNode) *ListNode {
	fast, slow := list, list
	for {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == nil {
			return slow
		}
		if fast.Next == nil {
			return slow.Next
		}
	}
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	sec := getPalindromeStart(head)
	r := reverseLinkedList(sec)
	// PrintAll(r)

	for r != nil {
		if r.Val != head.Val {
			return false
		}

		head = head.Next
		r = r.Next
	}

	return true
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
	r := isPalindrome2(nodeB)
	fmt.Println(r)

	nodeC5 := &ListNode{1, nil}
	nodeC4 := &ListNode{2, nodeC5}
	nodeC3 := &ListNode{3, nodeC4}
	nodeC2 := &ListNode{2, nodeC3}
	nodeC := &ListNode{1, nodeC2}

	rc := isPalindrome2(nodeC)
	fmt.Println(rc)

	nodeD2 := &ListNode{-129, nil}
	nodeD := &ListNode{-129, nodeD2}

	rd := isPalindrome2(nodeD)
	fmt.Println(rd)

	nodeE4 := &ListNode{1, nil}
	nodeE3 := &ListNode{2, nodeE4}
	nodeE2 := &ListNode{1, nodeE3}
	nodeE := &ListNode{1, nodeE2}

	re := isPalindrome2(nodeE)
	fmt.Println(re)

}
