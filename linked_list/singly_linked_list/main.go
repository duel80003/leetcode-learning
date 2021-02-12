package main

import "fmt"

type MyLinkedList struct {
	Size int
	Head *Node
}

type Node struct {
	Val  int
	Next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	var l MyLinkedList
	return l
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index >= this.Size {
		return -1
	}
	head := this.Head
	for i := 0; i < index; i++ {
		head = head.Next
	}
	return head.Val
}

/**AddAtHead Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.Size++
	this.Head = &Node{val, this.Head}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.Size == 0 {
		this.AddAtHead(val)
		return
	}
	cur := this.Head
	for cur.Next != nil {
		cur = cur.Next
	}

	this.Size++
	cur.Next = &Node{val, nil}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}
	if index == this.Size {
		this.AddAtTail(val)
		return
	}
	head := this.Head
	for i := 0; i < index; i++ {
		if i != index-1 {
			head = head.Next
		} else {
			head.Next = &Node{val, head.Next}
			this.Size++
			break
		}
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size {
		return
	}
	if index == 0 {
		this.Head = this.Head.Next
		this.Size--
		return
	}
	head := this.Head
	for i := 0; i < index; i++ {
		if i != index-1 {
			head = head.Next
		} else {
			head.Next = head.Next.Next
			this.Size--
			break
		}
	}
}

func (this *MyLinkedList) PrintAll() {
	head := this.Head
	for i := 0; i < this.Size; i++ {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	// obj := Constructor()
	// obj.AddAtHead(5)
	// obj.AddAtTail(6)
	// obj.AddAtIndex(2, 0)
	// obj.PrintAll()
	// fmt.Println("----")
	// fmt.Println(obj.Get(1))
	// fmt.Println("----")

	// obj.AddAtIndex(1, 7)
	// obj.AddAtIndex(1, 8)

	// fmt.Printf("size %v \n", obj.Size)
	// obj.PrintAll()
	// fmt.Println("----")

	// obj.DeleteAtIndex(1)
	// fmt.Printf("size %v \n", obj.Size)
	// obj.PrintAll()

	// fmt.Println("----")
	// obj.AddAtIndex(0, 9)
	// fmt.Printf("size %v \n", obj.Size)
	// obj.PrintAll()

	// test := Constructor()
	// test.AddAtHead(1)
	// test.AddAtTail(3)
	// test.AddAtIndex(1, 2)
	// fmt.Println(test.Get(1))
	// fmt.Println("----")
	// test.PrintAll()
	// fmt.Println("----")
	// test.DeleteAtIndex(1)
	// fmt.Println("----")
	// test.PrintAll()
	// fmt.Println("----")
	// fmt.Println(test.Get(1))

	// 	["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
	// [[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
	test := Constructor()
	test.AddAtHead(7)
	test.AddAtHead(2)
	test.AddAtHead(1)

	test.AddAtIndex(3, 0)

	test.DeleteAtIndex(2)
	test.AddAtHead(6)
	test.AddAtTail(4)

	fmt.Println(test.Get(4))
	test.AddAtHead(4)
	test.AddAtIndex(5, 0)
	test.AddAtHead(6)

	// 	["MyLinkedList","addAtTail","get"]
	// [[],[1],[0]]

	test1 := Constructor()
	test1.AddAtTail(1)
	fmt.Println(test1.Get(0))

}
