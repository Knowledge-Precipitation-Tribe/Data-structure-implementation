package RecursionOperation

import (
	"fmt"
	"go-implementation/tree/model"
)

//输入字符串来创建树，#代表空
func CreateTree(node *model.Node) {
	s := ""
	fmt.Println("请输入节点的值： ")
	fmt.Scanln(&s)
	if s != "#" {
		node.Val = s
		node.Left = &model.Node{}
		node.Right = &model.Node{}
		CreateTree(node.Left)
		CreateTree(node.Right)
	} else {
		node = nil
	}
}

//先序遍历显示树结构
func PreShowTree(node *model.Node) {
	if node != nil{
		fmt.Printf("%s %d\n", node.Val, len(node.Val))
		PreShowTree(node.Left)
		PreShowTree(node.Right)
	}
}

//中序遍历
func InShowTree(node *model.Node) {
	if node != nil {
		PreShowTree(node.Left)
		fmt.Printf("%s %d\n", node.Val, len(node.Val))
		PreShowTree(node.Right)
	}
}

//后序遍历
func PostShowTree(node *model.Node) {
	if node != nil {
		PreShowTree(node.Left)
		PreShowTree(node.Right)
		fmt.Printf("%s %d\n", node.Val, len(node.Val))
	}
}
