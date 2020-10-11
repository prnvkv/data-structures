package queue

import (
	"fmt"
)

// SIZE is Total Allowed Size of the Queue
const SIZE = 10

// Queue is basic structure that defines queue Data Structure
type Queue struct {
	front, rear, cap int
	queue            [SIZE]string
}

var q *Queue

// OpsOnQueue defines operations allowed on Queue Struct
type OpsOnQueue interface {
	EnQueue(item string)
	DeQueue() string
	IsFull() bool
	IsEmpty() bool
	Peek() (int, int)
	GetCapacity() int
	Display()
}

// NewQueue creates a new queue object
func NewQueue() OpsOnQueue {
	return &Queue{
		front: -1,
		rear:  -1,
		cap:   0,
		queue: [SIZE]string{},
	}
}

// EnQueue takes the item as an argument to add the item into the queue.
func (q *Queue) EnQueue(item string) {
	if q.IsFull() {
		fmt.Println("Queue is full")
		return
	}
	if q.front == -1 {
		q.front++
	}
	q.rear = q.rear + 1
	q.queue[q.rear] = item
	fmt.Printf("Successfully pushed the item %s \n", item)
	return
}

// DeQueue removes an item from the Queue and returns it.
func (q *Queue) DeQueue() string {
	if q.IsEmpty() {
		fmt.Println("Queue is empty")
		return ""
	}
	item := q.queue[q.front]
	q.front = q.front + 1
	return item
}

// IsFull checks if the Queue is full.
func (q *Queue) IsFull() bool {
	if (len(q.queue) == SIZE) || (q.rear == SIZE-1) {
		return false
	}
	return true
}

// IsEmpty checks if Queue is empty
func (q *Queue) IsEmpty() bool {
	if (q.front == -1) || len(q.queue) == 0 {
		return true
	}
	return false
}

// Peek returns the indicies of the front and rear pointers
func (q *Queue) Peek() (int, int) {
	return q.front, q.rear
}

// GetCapacity returns the current capacity of the queue
func (q *Queue) GetCapacity() int {
	return len(q.queue)
}

// Display prints the queue
func (q *Queue) Display() {
	fmt.Println("Displaying the elements in Queue...")
	for i := q.front; i <= q.rear; i++ {
		fmt.Println(q.queue[i])
	}
}
