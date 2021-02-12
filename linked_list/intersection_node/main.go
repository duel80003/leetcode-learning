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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := getLength(headA)
	l2 := getLength(headB)

	dif := 0
	var h1 *ListNode
	var h2 *ListNode

	if l1 > l2 {
		h1 = headA
		h2 = headB
		dif = l1 - l2
	} else {
		h1 = headB
		h2 = headA
		dif = l2 - l1
	}
	for i := 0; i < dif; i++ {
		h1 = h1.Next
	}

	for h1 != nil {
		if h1 != h2 {
			h1 = h1.Next
			h2 = h2.Next
		} else {
			return h1
		}
	}

	return nil
}

func main() {
	//listA: [4,1,8,4,5]
	//ListB: [5,6,1,8,4,5]
	//expected result: 8
	nodeC3 := &ListNode{5, nil}
	nodeC2 := &ListNode{4, nodeC3}
	nodeC1 := &ListNode{8, nodeC2}

	nodeA2 := &ListNode{5, nodeC1}
	nodeA := &ListNode{4, nodeA2}

	nodeB3 := &ListNode{1, nodeC1}
	nodeB2 := &ListNode{6, nodeB3}
	nodeB := &ListNode{5, nodeB2}
	fmt.Println(getIntersectionNode(nodeA, nodeB))

}
