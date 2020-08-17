# 反转链表

## 题目描述

反转一个单链表。

示例:

```go
输入: 1->2->3->4->5->NULL 
输出: 5->4->3->2->1->NULL 
```

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/reverse-linked-list](https://leetcode-cn.com/problems/reverse-linked-list) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil {
		tmp := fast.Next
		fast.Next = slow
		slow = fast
		fast = tmp
	}
	head.Next = nil
	return slow
}
```

