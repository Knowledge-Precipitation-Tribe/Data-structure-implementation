# 重建二叉树

## 题目描述

输入某二叉树的前序遍历和中序遍历的结果，请重建该二叉树。假设输入的前序遍历和中序遍历的结果中都不含重复的数字。

例如，给出

```text
前序遍历 preorder = [3,9,20,15,7] 
中序遍历 inorder = [9,3,15,20,7] 
```

返回如下的二叉树：

![](../../../.gitbook/assets/image%20%287%29.png)

限制：

`0 <= 节点个数 <= 5000`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof](https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
var indexForInOrders map[int]int

func buildTree(preorder []int, inorder []int) *TreeNode {
	indexForInOrders = make(map[int]int)
	for i := 0; i<len(inorder); i++{
		indexForInOrders[inorder[i]] = i
	}
	return reConstructBinaryTree(preorder, 0, len(preorder) - 1, 0)
}

func reConstructBinaryTree(preorder []int, preL int, preR int, inL int) *TreeNode {
	if preL > preR{
		return nil
	}
	root := &TreeNode{Val:preorder[preL]}
	inIndex := indexForInOrders[root.Val]
	leftTreeSize := inIndex - inL
	root.Left = reConstructBinaryTree(preorder, preL + 1, preL + leftTreeSize, inL)
	root.Right = reConstructBinaryTree(preorder, preL + leftTreeSize + 1, preR , inL + leftTreeSize + 1)
	return root
}
```

