# 三数之和

## 题目描述

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

```go
给定数组 nums = [-1, 0, 1, 2, -1, -4]，
满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]
```

{% hint style="info" %}
来源：力扣（LeetCode） 链接：[https://leetcode-cn.com/problems/3sum](https://leetcode-cn.com/problems/3sum) 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
{% endhint %}

## 题解

```go
import "sort"

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)

	result := [][]int{}
	for i := 0; i < length-2; i++ {
		for i != 0 && nums[i] == nums[i-1] && i < length-2 {
			i++
		}
		target := 0 - nums[i]
		l := i + 1
		r := length - 1
		for l < r {
			if nums[l]+nums[r] == target {
				result = append(result, []int{nums[i], nums[l], nums[r]})
				l++
				r--
				for nums[l] == nums[l-1] && nums[r] == nums[r+1] && l < r {
					l++
					r--
				}
			} else if nums[l]+nums[r] < target {
				l++
				for nums[l] == nums[l-1] && l < r {
					l++
				}
			} else {
				r--
				for nums[r] == nums[r+1] && l < r {
					r--
				}
			}
		}
	}
	return result
}
```

