package main

import (
	"fmt"
	"go-implementation/queue/sequential-queue"
)

func main() {
	queue := sequential_queue.Queue{}

	queue.InitQueue(3)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.EnQueue(4)

	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
	fmt.Println(queue.DeQueue())
}