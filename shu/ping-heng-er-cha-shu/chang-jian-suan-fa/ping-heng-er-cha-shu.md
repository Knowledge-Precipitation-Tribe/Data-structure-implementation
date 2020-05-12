# 平衡二叉树

## 题目描述

输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

![](../../../.gitbook/assets/image%20%2823%29.png)

![](../../../.gitbook/assets/image%20%2815%29.png)

限制：

`1 <= 树的结点个数 <= 10000`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof](https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func isBalanced(root *TreeNode) bool {
	return dfsIsBalanced(root) != -1
}

func dfsIsBalanced(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := dfsIsBalanced(root.Left)
	r := dfsIsBalanced(root.Right)
	if l != -1 && r != -1 && int(math.Abs(float64(l-r))) <= 1 {
		return int(math.Max(float64(l), float64(r))) + 1
	}
	return -1
}
```

