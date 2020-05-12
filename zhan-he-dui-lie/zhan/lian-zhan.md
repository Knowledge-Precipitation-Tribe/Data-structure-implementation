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

这里这里

```go
//数据出栈
func (stack *Stack)Pop() (interface{}, error){
	if stack.IsEmpty(){
		return nil, errors.New("empty stack")
	}
	v := stack.data[stack.top]
	stack.top--
	if stack.top != -1{
		stack.data = stack.data[:stack.top]
	}else{
		stack.InitStack()
	}
	return v, nil
}
```

## 完整实现

```go
type Stack struct {
	data []interface{}
	top int
}


//初始化栈
func (stack *Stack)InitStack(){
	stack.data = make([]interface{}, 0)
	stack.top = -1
}


//判断是否为空
func (stack *Stack)IsEmpty() bool{
	if stack.top == -1{
		return true
	}
	return false
}

//数据入栈
func (stack *Stack)Push(v interface{}){
	stack.top++
	stack.data = append(stack.data, v)
}

//数据出栈
func (stack *Stack)Pop() (interface{}, error){
	if stack.IsEmpty(){
		return nil, errors.New("empty stack")
	}
	v := stack.data[stack.top]
	stack.top--
	if stack.top != -1{
		stack.data = stack.data[:stack.top]
	}else{
		stack.InitStack()
	}
	return v, nil
}
```

