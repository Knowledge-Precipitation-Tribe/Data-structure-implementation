package LinkedListWithHead

import (
	"fmt"
	"go-implementation/model/LinkedList"
)

//创建带有头节点的链表
//使用头插法
func CreateHead(n int) *LinkedList.Node {
	head := &LinkedList.Node{
		Val:  -1,
		Next: nil,
	}
	p := head
	for i := 0; i < n; i++ {
		temp := &LinkedList.Node{
			Val:  i,
			Next: nil,
		}
		temp.Next = p.Next
		p.Next = temp
	}
	return head
}

//使用尾插法
func CreateTail(n int) *LinkedList.Node {
	head := &LinkedList.Node{
		Val:  -1,
		Next: nil,
	}
	p := head
	for i := 0; i < n; i++ {
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
	for temp := head.Next; temp != nil; temp = temp.Next {
		fmt.Printf("%d\t", temp.Val)
	}
	fmt.Println()
}
