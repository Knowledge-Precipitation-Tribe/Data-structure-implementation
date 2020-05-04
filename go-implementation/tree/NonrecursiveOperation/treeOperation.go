package NonrecursiveOperation

import (
	"fmt"
	"go-implementation/model/tree"
	"go-implementation/stack/sequential-stack"
)

//非递归先序遍历
func PreShow(node *tree.BTNode){
	if *node != (tree.BTNode{}) && node != nil{
		stack := &sequential_stack.Stack{}
		stack.InitStack()

		stack.Push(node)
		for !stack.IsEmpty(){
			temp, err := stack.Pop()
			btNode := temp.(*tree.BTNode)
			if err == nil{
				fmt.Println(btNode.Val)
				if *btNode.Left != (tree.BTNode{}) && btNode.Left != nil{
					stack.Push(btNode.Left)
				}
				if *btNode.Right != (tree.BTNode{}) && btNode.Right != nil{
					stack.Push(btNode.Right)
				}
			}else{
				return
			}
		}
	}
}

