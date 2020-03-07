package LinkedListWithoutHead

import (
	"../model"
	"fmt"
)

//不含头节点的链表
//头插法创建链表
func CreateHead(n int) *model.Node{
	head := &model.Node{
		Val:0,
		Next:nil,
	}
	for i := 1; i<n; i++{
		temp := &model.Node{
			Val:i,
			Next:nil,
		}
		temp.Next = head
		head = temp
	}
	return head
}

//尾插法创建链表
func CreateTail(n int) *model.Node{
	head := &model.Node{
		Val:0,
		Next:nil,
	}
	p := head
	for i := 1; i<n; i++{
		temp := &model.Node{
			Val:i,
			Next:nil,
		}
		p.Next = temp
		p = temp
	}
	return head
}

func ShowList(head *model.Node, name string)  {
	fmt.Printf("%s :\n", name)
	for temp := head; temp != nil; temp = temp.Next{
		fmt.Printf("%d\t", temp.Val)
	}
	fmt.Println()
}