package main

import "fmt"

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := &ListNode{-1, nil}
	dum := result
	carry := 0
	for l1 != nil && l2 != nil {
		tmp := &ListNode{}
		tmp.Val = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		dum.Next = tmp
		dum = dum.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	var remain *ListNode

	if l1 != nil {
		remain = l1
	} else if l2 != nil {
		remain = l2
	}
	for remain != nil {
		tmp := &ListNode{}
		tmp.Val = (remain.Val + carry) % 10
		carry = (remain.Val + carry) / 10
		dum.Next = tmp
		dum = dum.Next
		remain = remain.Next
	}
	// fmt.Printf("carry: %v \n", carry)
	if carry != 0 {
		dum.Next = &ListNode{carry, nil}
		dum = dum.Next
	}
	return result.Next
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

	l1_3 := &ListNode{3, nil}
	l1_2 := &ListNode{4, l1_3}
	l1_1 := &ListNode{2, l1_2}

	l2_3 := &ListNode{4, nil}
	l2_2 := &ListNode{6, l2_3}
	l2_1 := &ListNode{5, l2_2}

	result := addTwoNumbers(l1_1, l2_1)
	printAll(result)

	l1a7 := &ListNode{9, nil}
	l1a6 := &ListNode{9, l1a7}
	l1a5 := &ListNode{9, l1a6}
	l1a4 := &ListNode{9, l1a5}
	l1a3 := &ListNode{9, l1a4}
	l1a2 := &ListNode{9, l1a3}
	l1a1 := &ListNode{9, l1a2}

	l2a4 := &ListNode{9, nil}
	l2a3 := &ListNode{9, l2a4}
	l2a2 := &ListNode{9, l2a3}
	l2a1 := &ListNode{9, l2a2}
	result2 := addTwoNumbers(l1a1, l2a1)
	printAll(result2)

	l3c2 := &ListNode{7, nil}
	l3a1 := &ListNode{3, l3c2}

	l3b2 := &ListNode{2, nil}
	l3b1 := &ListNode{9, l3b2}

	result3 := addTwoNumbers(l3a1, l3b1)
	printAll(result3)
}
