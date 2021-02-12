package main

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			if fast == nil || fast.Next == nil {
				return nil
			}
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}

func main() {
	// // [3,2,0,-4 ]
	// node4 := &ListNode{-4, nil}
	// node3 := &ListNode{0, node4}
	// node2 := &ListNode{2, node3}
	// node4.Next = node2
	// head := &ListNode{3, node2}
	// fmt.Println(hasCycle(head))
	// fmt.Println(detectCycle(head))
	// //[]
	// head2 := &ListNode{}
	// fmt.Println(hasCycle(head2))
	// fmt.Println(detectCycle(head2))
	// //[1, 2]
	// head3 := &ListNode{1, nil}
	// node5 := &ListNode{2, head3}
	// head3.Next = node5
	// fmt.Println(hasCycle(head3))
	// fmt.Println(detectCycle(head3))

}
