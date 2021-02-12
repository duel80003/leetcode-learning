package main

import "fmt"

// ListNode is leetCode Learn data Structure
type ListNode struct {
	Val  int
	Next *ListNode
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

}
