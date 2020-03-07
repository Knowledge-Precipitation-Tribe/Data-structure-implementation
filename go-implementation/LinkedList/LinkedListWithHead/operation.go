package LinkedListWithHead

import (
	"../model"
	"fmt"
)

//创建带有头节点的链表
//使用头插法
func CreateHead(n int) *model.Node {
	head := &model.Node{
		Val:-1,
		Next:nil,
	}
	p := head
	for i := 0; i<n; i++{
		temp := &model.Node{
			Val:i,
			Next:nil,
		}
		temp.Next = p.Next
		p.Next = temp
	}
	return head
}

//使用尾插法
func CreateTail(n int) *model.Node {
	head := &model.Node{
		Val:-1,
		Next:nil,
	}
	p := head
	for i := 0; i<n; i++{
		temp := &model.Node{
			Val:i,
			Next:nil,
		}
		p.Next = temp
		p = temp
	}
	return head
}

func ShowList(head *model.Node, name string){
	fmt.Printf("%s :\n", name)
	for temp := head.Next; temp != nil; temp = temp.Next{
		fmt.Printf("%d\t", temp.Val)
	}
	fmt.Println()
}