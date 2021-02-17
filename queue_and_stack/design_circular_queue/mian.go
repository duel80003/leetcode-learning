package main

import "fmt"

// MyCircularQueue queue implemention
type MyCircularQueue struct {
	head  int
	tail  int
	count int
	arr   []int
}

// Constructor init MyCircularQueue instance
func Constructor(k int) MyCircularQueue {
	var q MyCircularQueue
	q.arr = make([]int, k)
	q.tail = -1
	return q
}

//EnQueue add element to tail
func (queue *MyCircularQueue) EnQueue(value int) bool {
	if queue.IsFull() {
		return false
	}
	queue.tail++
	if queue.tail == len(queue.arr) {
		queue.tail = 0
	}
	queue.arr[queue.tail] = value
	queue.count++
	return true
}

//DeQueue remote element from head
func (queue *MyCircularQueue) DeQueue() bool {
	if queue.IsEmpty() {
		return false
	}
	// fmt.Printf("before: head %v, tail %v \n", queue.head, queue.tail)
	queue.arr[queue.head] = 0
	queue.count--

	if queue.head == len(queue.arr)-1 {
		queue.head = 0
	} else if queue.count == 0 {
		queue.head, queue.tail = 0, -1
	} else {
		queue.head++
	}
	// fmt.Printf("head %v, tail %v \n", queue.head, queue.tail)
	return true
}

//Front return the head value
func (queue *MyCircularQueue) Front() int {
	if queue.count == 0 {
		return -1
	}
	return queue.arr[queue.head]
}

//Rear return the tail value
func (queue *MyCircularQueue) Rear() int {
	if queue.count == 0 {
		return -1
	}
	return queue.arr[queue.tail]
}

// IsEmpty return queue is emtpy
func (queue *MyCircularQueue) IsEmpty() bool {
	return queue.count == 0
}

// IsFull return queue is full
func (queue *MyCircularQueue) IsFull() bool {
	return queue.count == len(queue.arr)
}

func main() {
	// q1 := Constructor(3)
	// fmt.Println(q1.EnQueue(1))
	// fmt.Println(q1.EnQueue(2))
	// fmt.Println(q1.EnQueue(3))
	// fmt.Println(q1.arr)

	// fmt.Println(q1.EnQueue(1))
	// fmt.Println(q1.DeQueue())
	// fmt.Println(q1.DeQueue())
	// fmt.Println(q1.arr)
	// fmt.Println(q1.EnQueue(2))
	// fmt.Println(q1.arr)

	// ["MyCircularQueue","enQueue","deQueue","enQueue","enQueue","deQueue","isFull","isFull","Front","deQueue","enQueue","Front","enQueue","enQueue","Rear","Rear","deQueue","enQueue","enQueue","Rear","Rear","Front","Rear","Rear","deQueue","enQueue","Rear","deQueue","Rear","Rear","Front","Front","enQueue","enQueue","Front","enQueue","enQueue","enQueue","Front","isEmpty","enQueue","Rear","enQueue","Front","enQueue","enQueue","Front","enQueue","deQueue","deQueue","enQueue","deQueue","Front","enQueue","Rear","isEmpty","Front","enQueue","Front","deQueue","enQueue","enQueue","deQueue","deQueue","Front","Front","deQueue","isEmpty","enQueue","Rear","Front","enQueue","isEmpty","Front","Front","enQueue","enQueue","enQueue","Rear","Front","Front","enQueue","isEmpty","deQueue","enQueue","enQueue","Rear","deQueue","Rear","Front","enQueue","deQueue","Rear","Front","Rear","deQueue","Rear","Rear","enQueue","enQueue","Rear","enQueue"]
	// [[81],[69],[],[92],[12],[],[],[],[],[],[28],[],[13],[45],[],[],[],[24],[27],[],[],[],[],[],[],[88],[],[],[],[],[],[],[53],[39],[],[28],[66],[17],[],[],[47],[],[87],[],[92],[94],[],[59],[],[],[99],[],[],[84],[],[],[],[52],[],[],[86],[30],[],[],[],[],[],[],[45],[],[],[83],[],[],[],[22],[77],[23],[],[],[],[14],[],[],[90],[57],[],[],[],[],[34],[],[],[],[],[],[],[],[49],[59],[],[71]]

	// [null,true,true,true,true,true,false,false,12,true,true,28,true,true,45,45,true,true,true,27,27,13,27,27,true,true,88,true,88,88,24,24,true,true,24,true,true,true,24,false,true,47,true,24,true,true,24,true,true,true,true,true,53,true,84,false,53,true,53,true,true,true,true,true,66,66,true,false,true,45,17,true,false,17,17,true,true,true,23,17,17,true,false,true,true,true,57,true,57,87,true,true,34,92,34,true,34,34,true,true,59,true]
	// q2 := Constructor(81)
	// fmt.Print("null, ")
	// fmt.Print(q2.EnQueue(69), ", ")
	// fmt.Print(q2.DeQueue(), ", ")
	// fmt.Print(q2.EnQueue(92), ", ")
	// fmt.Print(q2.EnQueue(69), ", ")
	// fmt.Print(q2.DeQueue(), ", ")
	// fmt.Print(q2.DeQueue(), ", ")
	// fmt.Print(q2.IsFull(), ", ")
	// fmt.Print(q2.IsFull(), ", ")
	// fmt.Print(q2.Front(), ", ")
	// fmt.Print(q2.DeQueue(), ", ")
	// fmt.Print(q2.EnQueue(69), ", ")
	// fmt.Print(q2.Front(), ", ")
	// fmt.Print(q2.EnQueue(13), ", ")
	// fmt.Print(q2.EnQueue(45), ", ")
	// fmt.Print(q2.Rear(), ", ")
	// fmt.Print(q2.Rear(), ", ")
	// fmt.Print(q2.DeQueue(), ", ")
	// fmt.Print(q2.EnQueue(24), ", ")
	// fmt.Print(q2.EnQueue(27), ", ")
	// fmt.Print(q2.Rear(), ", ")
	// fmt.Print(q2.Rear(), ", ")
	// fmt.Print(q2.Front(), ", ")
	// fmt.Print(q2.Rear(), ", ")
	// fmt.Print(q2.Rear(), ", \n")
	// fmt.Println(q2.arr)

	q3 := Constructor(3)
	fmt.Print("null, ")
	fmt.Print(q3.EnQueue(2), ", ")
	fmt.Print(q3.Rear(), ", ")
	fmt.Print(q3.Front(), ", \n")
	fmt.Println(q3.arr)

}
