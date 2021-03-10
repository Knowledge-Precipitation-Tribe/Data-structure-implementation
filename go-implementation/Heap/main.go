package main

import (
	"fmt"
	"go-implementation/Heap/heap"
)

func main() {
	h := heap.NewHeap(10)
	h.Insert(13)
	h.Insert(2)
	h.Insert(1)
	h.Insert(9)
	fmt.Println(h)
}