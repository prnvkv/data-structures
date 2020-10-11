package main

import (
	"fmt"

	"github.com/prnvkv/data-structures/queue"
)

func main() {
	qq := queue.NewQueue()
	qq.EnQueue(readInput())
	qq.EnQueue(readInput())
	qq.EnQueue(readInput())
	qq.EnQueue(readInput())
	qq.Display()

	front, rear := qq.Peek()
	fmt.Println(front, rear)
	qq.DeQueue()
	front, rear = qq.Peek()
	fmt.Println(front, rear)
	qq.Display()
	fmt.Printf("Capacity is %d\n", qq.GetCapacity())
}

func readInput() string {
	var item string
	fmt.Scanf("%s", &item)
	return item
}
