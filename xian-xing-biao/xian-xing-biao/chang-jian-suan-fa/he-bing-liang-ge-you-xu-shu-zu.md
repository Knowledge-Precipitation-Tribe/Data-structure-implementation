# 合并两个有序数组

## 题目描述

给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

说明:

* 初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。 
* 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:

```go
输入: 
nums1 = [1,2,3,0,0,0], m = 3 
nums2 = [2,5,6], n = 3
输出: 
[1,2,2,3,5,6] 
```

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/merge-sorted-array](https://leetcode-cn.com/problems/merge-sorted-array) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
func merge(nums1 []int, m int, nums2 []int, n int) {
	k := m + n
	for i := k - 1; i >= 0; i-- {
		if n == 0 {
			break
		}
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}
}
```

