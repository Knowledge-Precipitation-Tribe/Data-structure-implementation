package main

import (
	"go-implementation/model/tree"
	"go-implementation/tree/RecursionOperation"
)

func main() {
	root := &tree.BTNode{}
	RecursionOperation.CreateTree(root)
	RecursionOperation.PreShowTree(root)
}
