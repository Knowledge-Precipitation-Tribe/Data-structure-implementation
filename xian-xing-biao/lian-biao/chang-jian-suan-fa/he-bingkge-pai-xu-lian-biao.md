# 合并K个排序链表

## 题目描述

合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:

```c
[ 
  1->4->5, 
  1->3->4, 
  2->6
] 
```

输出: 

```c
1->1->2->3->4->4->5->6
```

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/merge-k-sorted-lists](https://leetcode-cn.com/problems/merge-k-sorted-lists) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

使用分治算法解决此问题

```go
func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists) - 1)
}

func merge(lists []*ListNode, l int, r int) *ListNode{
	if l == r{
		return lists[l]
	}
	if l > r{
		return nil
	}
	mid := (l + r) >> 1
	return mergeTwoLists(merge(lists, l, mid), merge(lists, mid + 1, r))
}

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

