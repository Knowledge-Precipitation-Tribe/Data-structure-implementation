# 链队

## 链队的定义

![](../../.gitbook/assets/image%20%284%29.png)

```go
//用于存储队列中的数据
type QNode struct {
	data interface{}
	next *QNode
}

//用于标记队列的头和尾
type Queue struct {
	front *QNode
	rear *QNode
}
```

## 队列的状态

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

### 队满

在内存无限大的情况下，链队不存在栈满的情况

## 两个操作

### 入队

```go
//入队
func (queue *Queue)EnQueue(v interface{}){
	temp := &QNode{
		data:v,
		next:nil,
	}
	if queue.IsEmpty(){
		queue.front = temp
		queue.rear = temp
	}else{
		queue.rear.next = temp
		queue.rear = temp
	}
}
```

### 出队

```go
//出队
func (queue *Queue)DeQueue() (interface{}, error) {
	if queue.IsEmpty(){
		return  nil, errors.New("empty queue")
	}
	temp := queue.front
	//当队列中仅有一个元素时要将头尾指针置为空
	if queue.front == queue.rear{
		queue.front = nil
		queue.rear = nil
	}else{
		queue.front = queue.front.next
	}
	return temp.data, nil
}
```

## 完整实现

```go
//用于存储队列中的数据
type QNode struct {
	data interface{}
	next *QNode
}

//用于标记队列的头和尾
type Queue struct {
	front *QNode
	rear *QNode
}

func (queue *Queue)InitQueue(){

}

//判断对空
func (queue *Queue)IsEmpty() bool{
	if queue.rear == nil || queue.front == nil{
		return true
	}
	return false
}

//入队
func (queue *Queue)EnQueue(v interface{}){
	temp := &QNode{
		data:v,
		next:nil,
	}
	if queue.IsEmpty(){
		queue.front = temp
		queue.rear = temp
	}else{
		queue.rear.next = temp
		queue.rear = temp
	}
}

//出队
func (queue *Queue)DeQueue() (interface{}, error) {
	if queue.IsEmpty(){
		return  nil, errors.New("empty queue")
	}
	temp := queue.front
	//当队列中仅有一个元素时要将头尾指针置为空
	if queue.front == queue.rear{
		queue.front = nil
		queue.rear = nil
	}else{
		queue.front = queue.front.next
	}
	return temp.data, nil
}
```

