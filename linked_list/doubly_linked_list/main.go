package main

import "fmt"

// MyLinkedList is doubly linked list
type MyLinkedList struct {
	Size int
	Head *Node
}

// Node for linked list
type Node struct {
	Val  int
	Prev *Node
	Next *Node
}

// Constructor Initialize your data structure here.
func Constructor() MyLinkedList {
	var l MyLinkedList
	return l
}

// Get the value of the index-th node in the linked list. If the index is invalid, return -1.
func (list *MyLinkedList) Get(index int) int {
	if index >= list.Size {
		return -1
	}
	h := list.Head
	for i := 0; i < index; i++ {
		h = h.Next
	}
	fmt.Println(h)
	return h.Val
}

// AddAtHead add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list.
func (list *MyLinkedList) AddAtHead(val int) {
	if list.Size == 0 {
		list.Head = &Node{val, nil, nil}
		list.Size++
		return
	}
	newNode := &Node{val, nil, list.Head}
	list.Head.Prev = newNode
	list.Head = newNode
	list.Size++
}

// AddAtTail Append a node of value val to the last element of the linked list.
func (list *MyLinkedList) AddAtTail(val int) {
	if list.Size == 0 {
		list.Head = &Node{val, nil, nil}
		list.Size++
		return
	}
	h := list.Head
	target := list.Size - 1
	for i := 0; i < target; i++ {
		h = h.Next
	}
	newNode := &Node{val, h, nil}
	h.Next = newNode
	list.Size++
}

// AddAtIndex Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted.
func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index > list.Size {
		return
	}
	if index == 0 {
		list.AddAtHead(val)
		return
	}
	if index == list.Size {
		list.AddAtTail(val)
		return
	}
	h := list.Head
	target := index - 1
	for i := 0; i < target; i++ {
		h = h.Next
	}
	newNode := &Node{val, h, h.Next}
	next := h.Next
	next.Prev = newNode
	h.Next = newNode
	list.Size++

}

// DeleteAtIndex Delete the index-th node in the linked list, if the index is valid.
func (list *MyLinkedList) DeleteAtIndex(index int) {
	if index >= list.Size {
		return
	}
	if list.Size == 1 {
		list.Head = nil
		list.Size--
		return
	}

	h := list.Head
	for i := 0; i < index; i++ {
		h = h.Next
	}
	prev := h.Prev
	next := h.Next
	if prev != nil {
		prev.Next = next
	} else {
		list.Head = h.Next
		if list.Head != nil {
			list.Head.Prev = nil
		}
	}

	if next != nil {
		next.Prev = prev
	} else {
		if prev != nil {
			prev.Next = nil
		}
	}
	list.Size--
}

// PrintAll print all values in list
func (list *MyLinkedList) PrintAll() {
	head := list.Head
	fmt.Print("[ ")
	for i := 0; i < list.Size; i++ {
		fmt.Printf("%v", head.Val)
		head = head.Next
		if head != nil {
			fmt.Print(", ")
		}
	}
	fmt.Printf(" ] \n")
}

func main() {

	//  Your MyLinkedList object will be instantiated and called as such:
	// obj := Constructor()
	// obj.AddAtHead(1)
	// obj.AddAtTail(2)
	// obj.PrintAll()
	// val := obj.Get(1)
	// fmt.Printf("get value: %v \n", val)

	// obj.AddAtIndex(1, 7)
	// obj.PrintAll()
	// obj.DeleteAtIndex(2)
	// obj.PrintAll()

	// obj := Constructor()
	// obj.AddAtHead(1)
	// obj.PrintAll()
	// obj.AddAtTail(3)
	// obj.PrintAll()
	// obj.AddAtIndex(1, 2)
	// obj.PrintAll()
	// val := obj.Get(1)
	// fmt.Printf("get value: %v \n", val)
	// obj.DeleteAtIndex(2)
	// obj.PrintAll()
	// val = obj.Get(2)
	// fmt.Printf("get value: %v \n", val)

	// obj := Constructor()
	// obj.AddAtHead(1)
	// obj.DeleteAtIndex(0)
	// obj.PrintAll()

	// 	obj := Constructor()
	// 	obj.AddAtIndex(0, 10)
	// 	val := obj.Get(0)
	// 	fmt.Printf("get value: %v \n", val)
	// 	obj.AddAtIndex(0, 20)
	// 	obj.AddAtIndex(1, 30)
	// 	obj.PrintAll()
	obj := Constructor()
	obj.AddAtHead(2)
	obj.DeleteAtIndex(1)
	obj.AddAtHead(2)
	obj.AddAtHead(7)
	obj.AddAtHead(3)
	obj.AddAtHead(2)
	obj.AddAtHead(5)
	obj.AddAtTail(5)
	val := obj.Get(5)
	fmt.Printf("get value: %v \n", val)
	obj.PrintAll()
	obj.DeleteAtIndex(6)
	obj.DeleteAtIndex(4)

}
