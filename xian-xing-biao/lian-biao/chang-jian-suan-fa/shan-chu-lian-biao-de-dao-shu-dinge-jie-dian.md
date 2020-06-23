# 删除链表的倒数第N个节点

## 题目描述

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

```text
给定一个链表: 1->2->3->4->5, 和 n = 2.
当删除了倒数第二个节点后，链表变为 1->2->3->5. 
```

说明：

`给定的 n 保证是有效的。`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
  p := head
	q := head
	for ;n != 0;n--{
		p = p.Next
	}
	if p == nil{
		head = head.Next
	}else {
		for p.Next != nil{
			q = q.Next
			p = p.Next
		}
		q.Next = q.Next.Next
	}
	return head
}
```

