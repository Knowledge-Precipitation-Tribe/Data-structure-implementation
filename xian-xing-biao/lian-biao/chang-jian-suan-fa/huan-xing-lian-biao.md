# 环形链表

## 题目描述

给定一个链表，判断链表中是否有环。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

示例 1：

```go
输入：head = [3,2,0,-4], pos = 1 
输出：true 
解释：链表中有一个环，其尾部连接到第二个节点。
```

![](../../../.gitbook/assets/image%20%28154%29.png)

示例 2：

```go
输入：head = [1,2], pos = 0 
输出：true 
解释：链表中有一个环，其尾部连接到第一个节点。
```

![](../../../.gitbook/assets/image%20%28160%29.png)

示例 3：

```go
输入：head = [1], pos = -1 
输出：false 
解释：链表中没有环。
```

![](../../../.gitbook/assets/image%20%28159%29.png)

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/linked-list-cycle](https://leetcode-cn.com/problems/linked-list-cycle) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func hasCycle(head *ListNode) bool {
	if head == nil{
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil{
		if slow == fast{
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}
```

