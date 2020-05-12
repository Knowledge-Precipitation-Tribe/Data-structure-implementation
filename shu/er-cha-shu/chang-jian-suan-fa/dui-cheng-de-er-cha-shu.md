# 对称的二叉树

## 题目描述

请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

![](../../../.gitbook/assets/image%20%2819%29.png)

示例 1：

```text
输入：root = [1,2,2,3,4,4,3] 
输出：true 
```

示例 2：

```text
输入：root = [1,2,2,null,3,null,3] 
输出：false
```

限制：

`0 <= 节点个数 <= 1000`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof](https://leetcode-cn.com/problems/dui-cheng-de-er-cha-shu-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func isSymmetric(root *TreeNode) bool {
	return symmetric(root, root)
}

func symmetric(left *TreeNode, right *TreeNode) bool{
	if left == nil && right == nil{
		return true
	}else if left == nil || right == nil{
		return false
	}else if left.Val != right.Val{
		return false
	}
	return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}
```

