package linked_stack

import "errors"

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