# 移动零

## 题目描述

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

```text
输入: [0,1,0,3,12] 
输出: [1,3,12,0,0] 
```

说明:

必须在原数组上操作，不能拷贝额外的数组。 尽量减少操作次数。

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/move-zeroes](https://leetcode-cn.com/problems/move-zeroes) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func moveZeroes(nums []int)  {
	if len(nums) < 2 {
		return
	}

	last := -1
	for i, n := range nums {
		if n == 0 && last == -1 {
			last = i
		}
		if n != 0 && last != -1 {
			nums[last], nums[i] = nums[i], nums[last]
			last++
		}
	}
}
```

