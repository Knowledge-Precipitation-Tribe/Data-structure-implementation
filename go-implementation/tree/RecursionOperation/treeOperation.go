package RecursionOperation

import (
	"fmt"
	"go-implementation/model/tree"
)

//输入字符串来创建树，#代表空
func CreateTree(node *tree.BTNode) {
	s := ""
	fmt.Println("请输入节点的值： ")
	fmt.Scanln(&s)
	if s != "#" {
		node.Val = s
		node.Left = &tree.BTNode{}
		node.Right = &tree.BTNode{}
		CreateTree(node.Left)
		CreateTree(node.Right)
	} else {
		node = nil
	}
}

//先序遍历显示树结构
func PreShowTree(node *tree.BTNode) {
	if *node != (tree.BTNode{}) && node != nil{
		fmt.Printf("%s", node.Val)
		PreShowTree(node.Left)
		PreShowTree(node.Right)
	}
}

//中序遍历
func InShowTree(node *tree.BTNode) {
	if *node != (tree.BTNode{}) && node != nil{
		PreShowTree(node.Left)
		fmt.Printf("%s", node.Val)
		PreShowTree(node.Right)
	}
}

//后序遍历
func PostShowTree(node *tree.BTNode) {
	if *node != (tree.BTNode{}) && node != nil{
		PreShowTree(node.Left)
		PreShowTree(node.Right)
		fmt.Printf("%s", node.Val)
	}
}
