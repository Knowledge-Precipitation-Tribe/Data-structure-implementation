package main

import (
	"go-implementation/tree/RecursionOperation"
	"go-implementation/tree/model"
)

func main() {
	root := &model.Node{}
	RecursionOperation.CreateTree(root)
	RecursionOperation.PreShowTree(root)
}
