package LinkedListWithoutHead

import (
	"fmt"
	"go-implementation/model/LinkedList"
)

//不含头节点的链表
//头插法创建链表
func CreateHead(n int) *LinkedList.Node {
	head := &LinkedList.Node{
		Val:  0,
		Next: nil,
	}
	for i := 1; i < n; i++ {
		temp := &LinkedList.Node{
			Val:  i,
			Next: nil,
		}
		temp.Next = head
		head = temp
	}
	return head
}

//尾插法创建链表
func CreateTail(n int) *LinkedList.Node {
	head := &LinkedList.Node{
		Val:  0,
		Next: nil,
	}
	p := head
	for i := 1; i < n; i++ {
		temp := &LinkedList.Node{
			Val:  i,
			Next: nil,
		}
		p.Next = temp
		p = temp
	}
	return head
}

func ShowList(head *LinkedList.Node, name string) {
	fmt.Printf("%s :\n", name)
	for temp := head; temp != nil; temp = temp.Next {
		fmt.Printf("%d\t", temp.Val)
	}
	fmt.Println()
}
