# 树的子结构

## 题目描述

输入两棵二叉树A和B，判断B是不是A的子结构。\(约定空树不是任意一个树的子结构\)

B是A的子结构， 即 A中有出现和B相同的结构和节点值。

例如: 给定的树 A:

![](../../../.gitbook/assets/image%20%2829%29.png)

给定的树 B：

![](../../../.gitbook/assets/image%20%2831%29.png)

返回 true，因为 B 与 A 的一个子树拥有相同的结构和节点值。

示例 1：

```text
输入：A = [1,2,3], B = [3,1] 输出：false 
```

示例 2：

```text
输入：A = [3,4,5,1,2], B = [4,1] 输出：true
```

限制：

`0 <= 节点个数 <= 10000`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof](https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func isSubStructure(A *TreeNode, B *TreeNode) bool {
    if A == nil && B == nil {
        return true
    }
    if A == nil || B == nil{
        return false
    }

    var ret bool

    //当在 A 中找到 B 的根节点时，进入helper递归校验
    if A.Val == B.Val{
        ret = helper(A,B)
    }

    //ret == false，说明 B 的根节点不在当前 A 树顶中，进入 A 的左子树进行递归查找
    if !ret {
        ret = isSubStructure(A.Left,B)
    }

    //当 B 的根节点不在当前 A 树顶和左子树中，进入 A 的右子树进行递归查找
    if !ret {
        ret = isSubStructure(A.Right,B)
    }
    return ret
}

//helper 校验 B 是否与 A 的一个子树拥有相同的结构和节点值
func helper(a,b *TreeNode) bool{
    if b == nil{
        return true
    }
    if a == nil{
        return false
    }
    if a.Val != b.Val{
        return false
    }
    //a.Val == b.Val 递归校验 A B 左子树和右子树的结构和节点是否相同
    return helper(a.Left,b.Left) && helper(a.Right,b.Right)
}
```

