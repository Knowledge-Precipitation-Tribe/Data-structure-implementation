package main

import (
	"fmt"
	"go-implementation/stack/linked-stack"
)

func main() {
	//stack := &sequential_stack.Stack{}
    stack := &linked_stack.LNode{}

	stack.InitStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}