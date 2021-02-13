package main

import "fmt"

//ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
// 	result := &ListNode{-1, nil}
// 	prev := result
// 	for l1 != nil && l2 != nil {
// 		if l1.Val <= l2.Val {
// 			prev.Next = l1
// 			l1 = l1.Next
// 		} else {
// 			prev.Next = l2
// 			l2 = l2.Next
// 		}
// 		prev = prev.Next
// 	}
// 	if l1 == nil {
// 		prev.Next = l2
// 	} else if l2 == nil {
// 		prev.Next = l1
// 	}
// 	return result.Next
// }

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	result := &ListNode{-1, nil}
	dum := result
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			dum.Next = l1
			l1 = l1.Next
		} else {
			dum.Next = l2
			l2 = l2.Next
		}
		dum = dum.Next
	}
	if l1 == nil {
		dum.Next = l2
	} else if l2 == nil {
		dum.Next = l1
	}
	return result.Next
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
	l1_3 := &ListNode{4, nil}
	l1_2 := &ListNode{2, l1_3}
	l1_1 := &ListNode{1, l1_2}

	l2_3 := &ListNode{4, nil}
	l2_2 := &ListNode{3, l2_3}
	l2_1 := &ListNode{1, l2_2}

	result := mergeTwoLists(l1_1, l2_1)
	PrintAll(result)

	l1a1 := &ListNode{5, nil}
	l2b3 := &ListNode{4, nil}
	l2b2 := &ListNode{3, l2b3}
	l2b1 := &ListNode{1, l2b2}

	result2 := mergeTwoLists(l1a1, l2b1)
	PrintAll(result2)
}
