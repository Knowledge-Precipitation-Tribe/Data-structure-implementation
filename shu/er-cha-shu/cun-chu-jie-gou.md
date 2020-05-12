# 存储结构

## 顺序存储结构

顺序存储结构即用一个数组来存储一颗二叉树。

![](../../.gitbook/assets/image%20%2812%29.png)

例如此时知道根结点A的编号是$$i=0$$，那么它的左孩子结点编号就是 2 \* i+1=1 ，右孩子结点编号就是 2\*i+2=2。

{% hint style="info" %}
存储一般的二叉树会非常浪费空间，最适合存储完全二叉树
{% endhint %}

## 链式存储结构

观察二叉树的结构，用顺序存储显然不能很好的表示一颗二叉树，故使用带有两个指针域一个数据域的结构来创建二叉树的链式存储结构，如图所示。

![](../../.gitbook/assets/image%20%2811%29.png)

其中data用于存储当前结点的数据，left指针用于存储其左孩子的位置，right指针用于存储其右孩子的位置。在golang中定义如下。

```text
type TreeNode struct {
	Data string
	Left *TreeNode
	Right *TreeNode
}
```

![](../../.gitbook/assets/image%20%284%29.png)

