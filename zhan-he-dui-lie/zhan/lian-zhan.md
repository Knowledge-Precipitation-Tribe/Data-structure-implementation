# 链栈

## 链栈的定义

```go
type LNode struct {
	data interface{}
	next *LNode
}
```

## 栈的状态

### 栈空

```go
node.next == nil
```

```go
// 判断栈空
func (node *LNode)IsEmpty() bool{
	if node.next == nil{
		return true
	}
	return false
}
```

### 栈满

在内存无限大的情况下，链栈不存在栈满的情况

### 非法情况

* **下溢：**栈空仍然继续出栈

## 两个操作

### 入栈

这里就是含头结点的链表的插入

```go
// 入栈
func (node *LNode)Push(v interface{}){
	temp := &LNode{
		data:v,
		next:nil,
	}
	temp.next = node.next
	node.next = temp
}
```

### 出栈

这里就是含头结点的删除

```go
// 出栈
func (node *LNode)Pop() (interface{}, error) {
	if node.IsEmpty(){
		return nil, errors.New("empty stack")
	}
	var p *LNode
	p = node.next
	node.next = p.next
	return p.data, nil
}
```

## 完整实现

```go
type LNode struct {
	data interface{}
	next *LNode
}

func (node *LNode)InitStack(){

}

// 判断栈空
func (node *LNode)IsEmpty() bool{
	if node.next == nil{
		return true
	}
	return false
}

// 入栈
func (node *LNode)Push(v interface{}){
	temp := &LNode{
		data:v,
		next:nil,
	}
	temp.next = node.next
	node.next = temp
}

// 出栈
func (node *LNode)Pop() (interface{}, error) {
	if node.IsEmpty(){
		return nil, errors.New("empty stack")
	}
	var p *LNode
	p = node.next
	node.next = p.next
	return p.data, nil
}
```

