# 二叉搜索树的第k大节点

## 题目描述

给定一棵二叉搜索树，请找出其中第k大的节点。

![](../../../.gitbook/assets/image%20%2814%29.png)

![](../../../.gitbook/assets/image%20%2822%29.png)

限制：

`1 ≤ k ≤ 二叉搜索树元素个数`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof](https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

![](../../../.gitbook/assets/image%20%2825%29.png)

```go
func kthLargest(root *TreeNode, k int) int {
    ret := []int{}
    inOrder(root, &ret)
    return ret[k-1]
}

func inOrder(root *TreeNode, ret *[]int) {
    if root.Right != nil {
        inOrder(root.Right, ret)
    }
    if root != nil {
        *ret = append(*ret, root.Val)
    }
    if root.Left != nil {
        inOrder(root.Left, ret)
    }
}
```

