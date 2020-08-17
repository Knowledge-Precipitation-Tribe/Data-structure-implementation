# 盛最多水的容器

## 题目描述

给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 \(i, ai\) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 \(i, ai\) 和 \(i, 0\)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。

![](../../../.gitbook/assets/image%20%28156%29.png)

图中垂直线代表输入数组 \[1,8,6,2,5,4,8,3,7\]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

示例：

```go
输入：[1,8,6,2,5,4,8,3,7] 
输出：49
```

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/container-with-most-water](https://leetcode-cn.com/problems/container-with-most-water) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func maxArea(height []int) int {
	if len(height) == 0{
		return 0
	}
	i := 0
	j := len(height) - 1
	ans := 0
	for i < j{
		area := min(height[i], height[j]) * (j - i)
		ans = max(ans, area)
		if height[i] <= height[j]{
			i += 1
		}else{
			j -= 1
		}
	}
	return ans
}

func min(a, b int) int{
	if a < b{
		return a
	}
	return b
}

func max(a, b int) int{
	if a > b{
		return a
	}
	return b
}
```

