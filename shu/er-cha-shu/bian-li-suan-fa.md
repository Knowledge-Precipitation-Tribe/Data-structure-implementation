# 遍历算法

## 先序遍历

先序遍历算法流程如下：

1. 访问根结点
2. 先序遍历左子树
3. 先序遍历右子树

{% tabs %}
{% tab title="递归方式" %}
```go
func Preorder(node *TreeNode){
	if *node != (tree.BTNode{}) && node != nil{
		fmt.Printf("%s", node.Val)
		PreShowTree(node.Left)
		PreShowTree(node.Right)
	}
}
```
{% endtab %}

{% tab title="非递归方式" %}
```
func PreShow(node *tree.BTNode){
	if *node != (tree.BTNode{}) && node != nil{
		stack := &sequential_stack.Stack{}
		stack.InitStack()

		stack.Push(node)
		for !stack.IsEmpty(){
			temp, err := stack.Pop()
			btNode := temp.(*tree.BTNode)
			if err == nil{
				fmt.Println(btNode.Val)
				if *btNode.Left != (tree.BTNode{}) && btNode.Left != nil{
					stack.Push(btNode.Left)
				}
				if *btNode.Right != (tree.BTNode{}) && btNode.Right != nil{
					stack.Push(btNode.Right)
				}
			}
		}
	}
}
```
{% endtab %}
{% endtabs %}

## 中序遍历

中序遍历算法流程如下：

1. 中序遍历左子树
2. 访问根结点
3. 中序遍历右子树

{% tabs %}
{% tab title="递归方式" %}
```go
func InShowTree(node *model.Node){
	if *node != (tree.BTNode{}) && node != nil{
		PreShowTree(node.Left)
		fmt.Printf("%s", node.Val)
		PreShowTree(node.Right)
	}
}
```
{% endtab %}

{% tab title="非递归方式" %}
```

```
{% endtab %}
{% endtabs %}

## 后序遍历

后序遍历算法流程如下：

1. 后后序遍历左子树
2. 后序遍历右子树
3. 访问根结点

{% tabs %}
{% tab title="递归方式" %}
```go
func PostShowTree(node *model.Node){
	if *node != (tree.BTNode{}) && node != nil{
		PreShowTree(node.Left)
		PreShowTree(node.Right)
		fmt.Printf("%s", node.Val)
	}
}
```
{% endtab %}

{% tab title="非递归方式" %}
```

```
{% endtab %}
{% endtabs %}

## 层次遍历



{% tabs %}
{% tab title="递归方式" %}
```go
var leveles = make([][]int, 0)

func levelOrder1(root *TreeNode) [][]int {
	if root == nil{
		return leveles
	}
	helper(root, 0)
	return leveles
}

func helper(node *TreeNode, level int){
	if len(leveles) == level{
		leveles = append(leveles, make([]int, 0))
	}

	leveles[level] = append(leveles[level], node.Val)

	if node.Left != nil{
		helper(node.Left, level + 1)
	}
	if node.Right != nil{
		helper(node.Right, level + 1)
	}
}
```
{% endtab %}

{% tab title="非递归方式" %}
```go
func levelOrder(root *TreeNode) [][]int {
	leveles = make([][]int , 0)

	if root == nil{
		return leveles
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) != 0{
		leveles = append(leveles, make([]int ,0))
		level_length := len(queue)
		for i := 0; i < level_length; i++{
			node := queue[0]
			queue = queue[1:]
			leveles[level] = append(leveles[level], node.Val)
			if node.Left != nil{
				queue = append(queue, node.Left)
			}
			if node.Right != nil{
				queue = append(queue, node.Right)
			}
		}
		level++
	}
	return leveles
}
```
{% endtab %}
{% endtabs %}

