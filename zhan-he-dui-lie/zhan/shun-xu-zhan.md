# 顺序栈

## 顺序栈的定义

```go
type Stack struct {
	data []interface{}
	// 定位栈顶未知
	top int
}
```

## 栈的状态

### 栈空

```go
top == -1
```

当然`top == 0`也是可以的，只不过那样会浪费一个存储空间。

```go
//判断是否为空
func (stack *Stack)IsEmpty() bool{
	if stack.top == -1{
		return true
	}
	return false
}
```

### 栈满

```go
top == MAxSize - 1
```

其实一般来说会给栈规定一个最大存储容量MaxSize，但是在go中我们可以使用切片方式理论上来实现无限栈。

### 非法情况

* **上溢：**栈满仍然继续入栈
* **下溢：**栈空仍然继续出栈

## 两个操作

### 入栈

注意我们的这种实现方式要先移动栈顶指针top，在执行入队操作。

```go
//数据入栈
func (stack *Stack)Push(v interface{}){
	stack.top++
	stack.data = append(stack.data, v)
}
```

### 出栈

注意我们的这种实现方式要取出元素，再移动栈顶指针top。

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

