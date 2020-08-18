# 最小栈

## 题目描述

设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

push\(x\) —— 将元素 x 推入栈中。 pop\(\) —— 删除栈顶的元素。 top\(\) —— 获取栈顶元素。 getMin\(\) —— 检索栈中的最小元素。

示例:

```go
输入： 
["MinStack","push","push","push","getMin","pop","top","getMin"] 
[[],[-2],[0],[-3],[],[],[],[]]
输出： 
[null,null,null,null,-3,null,0,-2]
解释： 
MinStack minStack = new MinStack(); 
minStack.push(-2); 
minStack.push(0); 
minStack.push(-3); 
minStack.getMin(); --> 返回 -3. 
minStack.pop(); 
minStack.top(); --> 返回 0. 
minStack.getMin(); --> 返回 -2.
```

提示：

* pop、top 和 getMin 操作总是在非空栈上调用。

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/min-stack](https://leetcode-cn.com/problems/min-stack) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
type MinStack struct {
	stack []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack, x)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(x, top))
}

func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
```

