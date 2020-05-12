package sequential_queue

import (
	"errors"
)

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