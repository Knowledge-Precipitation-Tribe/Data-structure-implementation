package main

import (
	"fmt"
	"go-implementation/model/tree"
	"go-implementation/tree/NonrecursiveOperation"
	"go-implementation/tree/RecursionOperation"
)

func main() {
	root := &tree.BTNode{}
	RecursionOperation.CreateTree(root)
	//value := root
	RecursionOperation.PreShowTree(root)
	fmt.Println("-------")
	NonrecursiveOperation.PreShow(root)
}
