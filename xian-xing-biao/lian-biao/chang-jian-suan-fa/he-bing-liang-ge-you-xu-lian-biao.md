# 合并两个有序链表

## 题目描述

将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

```text
输入：1->2->4, 1->3->4 
输出：1->1->2->3->4->4
```

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/merge-two-sorted-lists](https://leetcode-cn.com/problems/merge-two-sorted-lists) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
  head := &ListNode{}
	
	p := head
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			p.Next = l1
			l1 = l1.Next
		}else{
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 == nil{
		p.Next = l2
	}else{
		p.Next = l1
	}
	return head.Next
}
```



