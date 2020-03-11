package main

import (
	"../RecursionOperation"
	"../model"
)

func main() {
	root := &model.Node{}
	RecursionOperation.CreateTree(root)
	RecursionOperation.PreShowTree(root)

}