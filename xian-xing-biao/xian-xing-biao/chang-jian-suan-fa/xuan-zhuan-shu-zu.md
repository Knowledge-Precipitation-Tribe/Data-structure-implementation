# 旋转数组

## 题目描述

给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

示例 1:

```go
输入: [1,2,3,4,5,6,7] 和 k = 3 
输出: [5,6,7,1,2,3,4] 
解释: 
向右旋转 1 步: [7,1,2,3,4,5,6] 
向右旋转 2 步: [6,7,1,2,3,4,5] 
向右旋转 3 步: [5,6,7,1,2,3,4] 
```

示例 2:

```go
输入: [-1,-100,3,99] 和 k = 2 
输出: [3,99,-1,-100] 
解释: 
向右旋转 1 步: [99,-1,-100,3] 
向右旋转 2 步: [3,99,-1,-100] 
```

说明:

* 尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。 
* 要求使用空间复杂度为 O\(1\) 的 原地 算法。

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/rotate-array](https://leetcode-cn.com/problems/rotate-array) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func rotate(nums []int, k int)  {
	if nums == nil || len(nums) == 0{
		return
	}
	length := len(nums)
	k %= k % length
	reverse(nums, 0, length - k - 1)
	
	reverse(nums, length - k - 1, length - 1)
	reverse(nums, 0 , length - 1)
}

func reverse(nums []int, i, j int){
	for i < j{
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
```

