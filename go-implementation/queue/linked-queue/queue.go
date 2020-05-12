package linked_queue

import "errors"

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