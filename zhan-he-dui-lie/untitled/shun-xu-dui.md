# 顺序队

## 循环队列

在顺序队中，通常让队尾指针rear指向刚进队的元素位置，让队首指针front指向刚出队的元素位置，因此，元素进队的时候，rear要向后移动；元素出队的时候，front也要向后移动。这样经过一些列的出队和进队操作后，两个指针最终会到达数组末端的`maxSize-1`处。这时虽然队列中已经没有元素了，但仍然无法让元素进队，这就是所谓的**假溢出**现象。

为了解决这个问题，我们采取循环队列的方式，就是将数组弄成逻辑上的环状。让front和rear沿着环走，这样就不会出现假溢出现象。

![](../../.gitbook/assets/image%20%2825%29.png)

从图示中可以看到，我们必须牺牲一个空间，以此来判断队列是否为空。

## 顺序队的定义

```go
type Queue struct {
	data []interface{}
	maxSize int
	front int
	rear int
}
```

* maxSize用于存储队列能存储的最多元素数目
* front代表队头
* rear代表队尾

## 队列的状态

### 队空

```go
queue.front == queue.rear
```

```go
//判断队空
func (queue *Queue)IsEmpty() bool{
	if queue.front == queue.rear{
		return true
	}
	return false
}
```

### 队满

```go
(queue.rear + 1) % queue.maxSize == queue.front
```

## 两个操作

### 入队

移动队尾指针

```go
//入队
func (queue *Queue)EnQueue(v interface{}) error{
	if queue.IsFull(){
		return errors.New("full queue")
	}
	queue.rear = (queue.rear + 1) % queue.maxSize
	queue.data[queue.rear] = v
	return nil
}
```

### 出队

移动队头指针

```go
//出队
func (queue *Queue)DeQueue() (interface{}, error) {
	if queue.IsEmpty(){
		return nil, errors.New("empty queue")
	}
	queue.front = (queue.front + 1) % queue.maxSize
	return queue.data[queue.front], nil
}
```

## 完整实现

```go
type Queue struct {
	data []interface{}
	maxSize int
	front int
	rear int
}

//初始化队列
func (queue *Queue)InitQueue(maxSize int){
	queue.data = make([]interface{}, maxSize)
	queue.maxSize = maxSize
	queue.front = 0
	queue.rear = 0
}

//判断队空
func (queue *Queue)IsEmpty() bool{
	if queue.front == queue.rear{
		return true
	}
	return false
}

//判断队满
func (queue *Queue)IsFull() bool{
	if (queue.rear + 1) % queue.maxSize == queue.front{
		return true
	}
	return false
}

//入队
func (queue *Queue)EnQueue(v interface{}) error{
	if queue.IsFull(){
		return errors.New("full queue")
	}
	queue.rear = (queue.rear + 1) % queue.maxSize
	queue.data[queue.rear] = v
	return nil
}

//出队
func (queue *Queue)DeQueue() (interface{}, error) {
	if queue.IsEmpty(){
		return nil, errors.New("empty queue")
	}
	queue.front = (queue.front + 1) % queue.maxSize
	return queue.data[queue.front], nil
}
```

