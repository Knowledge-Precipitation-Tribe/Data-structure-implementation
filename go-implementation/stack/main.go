package main

import (
	"fmt"
	"go-implementation/stack/sequential-stack"
)

func main() {
	stack := &sequential_stack.Stack{}

	stack.InitStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}