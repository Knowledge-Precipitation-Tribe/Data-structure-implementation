# 两两交换链表中的节点

## 题目描述

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

```go
给定 1->2->3->4, 你应该返回 2->1->4->3. 
```

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/swap-nodes-in-pairs](https://leetcode-cn.com/problems/swap-nodes-in-pairs) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	newHead := &ListNode{}
	newHead.Next = head
	h := newHead
	p := head
	q := head.Next
	for q != nil{
		p.Next = q.Next
		q.Next = p
		h.Next = q
		h = p
		p = p.Next
		if p == nil{
			break
		}
		q = p.Next
	}
	return newHead.Next
}
```

