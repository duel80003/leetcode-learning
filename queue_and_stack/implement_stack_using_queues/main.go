package main

import "fmt"

// MyStack instacke for stack implementation
type MyStack struct {
	EnQueue []int
	DeQueue []int
}

// Constructor initialize your data structure here.
func Constructor() MyStack {
	return MyStack{
		[]int{},
		[]int{},
	}
}

// Push element x onto stack.
func (stack *MyStack) Push(x int) {
	stack.EnQueue = append(stack.EnQueue, x)
}

//Pop the element on top of the stack and returns that element.
func (stack *MyStack) Pop() int {
	length := len(stack.EnQueue) - 1
	for i := 0; i < length; i++ {
		fmt.Println("aaa")
		stack.DeQueue = append(stack.DeQueue, stack.EnQueue[0])
		stack.EnQueue = stack.EnQueue[1:]
	}
	fmt.Println(stack.EnQueue)
	top := stack.EnQueue[0]
	stack.EnQueue = stack.DeQueue
	stack.DeQueue = nil
	return top
}

// Top return the top element.
func (stack *MyStack) Top() int {
	top := stack.Pop()
	stack.EnQueue = append(stack.EnQueue, top)
	return top
}

// Empty returns whether the stack is empty.
func (stack *MyStack) Empty() bool {
	return len(stack.EnQueue) == 0
}

func main() {
	// s1 := Constructor()

	// s1.Push(1)
	// fmt.Println(s1.Top())
	// fmt.Println(s1.Pop())
	// fmt.Println(s1.Empty())

	// s2 := Constructor()

	// s2.Push(1)
	// s2.Push(2)

	// fmt.Println(s2.Top())
	// s2.Push(3)

	// fmt.Println(s2.Top())

	s3 := Constructor()

	s3.Push(1)
	s3.Push(2)
	s3.Push(3)
	fmt.Println(s3.Pop())
	// fmt.Println(s3.Pop())
	// fmt.Println(s3.Pop())
	// fmt.Println(s3.Empty())
}
