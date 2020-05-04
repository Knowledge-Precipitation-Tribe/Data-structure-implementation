package main

import (
	"go-implementation/LinkedList/LinkedListWithoutHead"
)

func main() {
	//head := LinkedListWithHead.CreateHead(10)
	//LinkedListWithHead.ShowList(head, "l1")
	//
	//head2 := LinkedListWithHead.CreateTail(10)
	//LinkedListWithHead.ShowList(head2, "l2")

	head := LinkedListWithoutHead.CreateHead(10)
	LinkedListWithoutHead.ShowList(head, "l1")

	head2 := LinkedListWithoutHead.CreateTail(10)
	LinkedListWithoutHead.ShowList(head2, "l2")
}