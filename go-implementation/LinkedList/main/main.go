package main

import "../LinkedListWithHead"

func main() {
	head := LinkedListWithHead.CreateHead(10)
	LinkedListWithHead.ShowList(head, "l1")

	head2 := LinkedListWithHead.CreateTail(10)
	LinkedListWithHead.ShowList(head2, "l2")
}