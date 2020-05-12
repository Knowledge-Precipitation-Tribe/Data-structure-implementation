# 从上倒下打印二叉树/层次遍历

## 题目描述

从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

![](../../../.gitbook/assets/image%20%282%29.png)

提示：

`节点总数 <= 1000`

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof](https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func levelOrder(root *TreeNode) []int {
  if root == nil{
		return nil
	}
	buff := make([]*TreeNode, 0)
	buff = append(buff, root)
	result := make([]int, 0)
	for len(buff) != 0{
		temp := buff[0]
		if temp != nil{
			result = append(result, temp.Val)
            buff = append(buff, temp.Left)
		    buff = append(buff, temp.Right)
		}
		buff = buff[1:]
	}
	return result
}
```

