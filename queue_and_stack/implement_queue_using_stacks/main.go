package main

import "fmt"

// MyQueue instacke for queue implementation
type MyQueue struct {
	pushStack []int
	popStack  []int
}

// Constructor Initialize your data structure here.
func Constructor() MyQueue {

	return MyQueue{
		[]int{},
		[]int{},
	}
}

// Push element x to the back of queue.
func (queue *MyQueue) Push(x int) {
	queue.pushStack = append(queue.pushStack, x)
}

// Pop the element from in front of queue and returns that element.
func (queue *MyQueue) Pop() int {
	if len(queue.popStack) == 0 {
		queue.Peek()
	}
	val := queue.popStack[len(queue.popStack)-1]
	queue.popStack = queue.popStack[:len(queue.popStack)-1]
	return val
}

// Peek get the front element.
func (queue *MyQueue) Peek() int {
	if len(queue.popStack) == 0 {
		for len(queue.pushStack) != 0 {
			val := queue.pushStack[len(queue.pushStack)-1]
			queue.pushStack = queue.pushStack[:len(queue.pushStack)-1]
			queue.popStack = append(queue.popStack, val)
		}
	}

	return queue.popStack[len(queue.popStack)-1]
}

// Empty returns whether the queue is empty.
func (queue *MyQueue) Empty() bool {
	return len(queue.popStack) == 0 && len(queue.pushStack) == 0
}

func main() {
	q1 := Constructor()
	q1.Push(1)
	fmt.Println(q1.Peek())

	// ["MyCircularQueue","Push","Pop","Push","Push","Pop","Pop","Empty","Push","Peek","Push","Push","Pop","Push","Push","Peek"
	// null, 69, 92, 69, true, 69, 69, 13
	q2 := Constructor()
	fmt.Print("null, ")
	q2.Push(69)
	fmt.Print(q2.Pop(), ", ")
	q2.Push(92)
	q2.Push(69)

	fmt.Print(q2.Pop(), ", ")

	fmt.Print(q2.Pop(), ", ")
	fmt.Print(q2.Empty(), ", ")
	q2.Push(69)
	fmt.Print(q2.Peek(), ", ")
	q2.Push(13)
	q2.Push(45)
	fmt.Print(q2.Pop(), ", ")
	q2.Push(24)
	q2.Push(27)
	fmt.Print(q2.Peek(), ", ")
	// fmt.Println(q2)

}
