package sequential_stack

import "errors"

type Stack struct {
	data []interface{}
	top int
}


//初始化栈
func (stack *Stack)InitStack(){
	stack.data = make([]interface{}, 0)
	stack.top = -1
}

func (stack *Stack)IsEmpty() bool{
	if stack.top == -1{
		return true
	}
	return false
}

func (stack *Stack)Push(v interface{}){
	stack.top++
	stack.data = append(stack.data, v)
}

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