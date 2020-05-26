# 带头结点单链表

## 单链表定义

```go
type Node struct{
	Val int
	Next *Node
}
```

## 常用操作

### 头插法建立链表

```go
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
```

### 尾插法建立链表

```go
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
```



### 显示链表内容

```go
func ShowList(head *LinkedList.Node, name string) {
	fmt.Printf("%s :\n", name)
	for temp := head.Next; temp != nil; temp = temp.Next {
		fmt.Printf("%d\t", temp.Val)
	}
	fmt.Println()
}
```

