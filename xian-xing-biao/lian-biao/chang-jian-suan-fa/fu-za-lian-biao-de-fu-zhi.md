# 复杂链表的复制

## 题目描述

请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，还有一个 random 指针指向链表中的任意节点或者 null。

示例 1：

![](../../../.gitbook/assets/image%20%28104%29.png)

```text
输入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]] 
输出：[[7,null],[13,0],[11,4],[10,2],[1,0]] 
```

示例 2：

![](../../../.gitbook/assets/image%20%2895%29.png)

```text
输入：head = [[1,1],[2,1]] 
输出：[[1,1],[2,1]] 
```

示例 3：

![](../../../.gitbook/assets/image%20%2886%29.png)

```text
输入：head = [[3,null],[3,0],[3,null]] 
输出：[[3,null],[3,0],[3,null]] 
```

示例 4：

输入：head = \[\] 

输出：\[\] 

解释：给定的链表为空（空指针），因此返回 null。

提示：

* -10000 &lt;= Node.val &lt;= 10000 
* Node.random 为空（null）或指向链表中的节点。 
* 节点数目不超过 1000 。

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof](https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func copyRandomList(head *Node) *Node {
	var nodeMap = make(map[*Node]*Node)

	return copyDfs(head, nodeMap)
}

func copyDfs(head *Node, nodeMap map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}

	if v, ok := nodeMap[head]; ok {
		return v
	}
	copy := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	nodeMap[head] = copy
	copy.Next = copyDfs(head.Next, nodeMap)
	copy.Random = copyDfs(head.Random, nodeMap)
	return copy
}
```

