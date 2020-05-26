# 二叉树中和为某一值的路径

## 题目描述

输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

示例: 给定如下二叉树，以及目标和 `sum = 22`，

![](../../../.gitbook/assets/image%20%2842%29.png)

返回:

```go
[ 
    [5,4,11,2], 
    [5,8,4,5] 
]
```

提示：

`节点总数 <= 10000`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof](https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	dfs(root,sum,[]int{},&ret)
	return ret
}


func dfs(root *TreeNode,sum int,arr []int,ret *[][]int){
	if root == nil{
		return
	}
	arr = append(arr,root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil {
		//slice是一个指向底层的数组的指针结构体
		//因为是先序遍历，如果 root.Right != nil ,arr 切片底层的数组会被修改
		//所以这里需要 copy arr 到 tmp，再添加进 ret，防止 arr 底层数据修改带来的错误
		tmp := make([]int,len(arr))
		copy(tmp,arr)
		*ret = append(*ret,tmp)
	}

	dfs(root.Left,sum - root.Val,arr,ret)
	dfs(root.Right,sum - root.Val,arr,ret)

	arr = arr[:len(arr)-1]
}
```

