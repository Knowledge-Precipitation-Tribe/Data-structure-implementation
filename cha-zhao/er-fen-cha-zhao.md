# 二分查找

## 什么是二分查找

在计算机科学中，二分搜索（binary search），也称折半搜索（half-interval search）、对数搜索（logarithmic search），是一种在**有序数组**中查找某一特定元素的搜索算法。搜索过程从数组的中间元素开始，如果中间元素正好是要查找的元素，则搜索过程结束；如果某一特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间元素开始比较。如果在某一步骤数组为空，则代表找不到。这种搜索算法每一次比较都使搜索范围缩小一半。

![](../.gitbook/assets/image%20%2898%29.png)

## 二分查找时间复杂度分析

我们假设数据大小是 n，每次查找后数据都会缩小为原来的一半，也就是会除以 2。最坏情况下，直到查找区间被缩小为空，才停止。

![](../.gitbook/assets/image%20%2883%29.png)

可以看出来，这是一个等比数列。其中 n/2k=1 时，k 的值就是总共缩小的次数。而每一次缩小操作只涉及两个数据的大小比较，所以，经过了 k 次区间缩小操作，时间复杂度就是 O\(k\)。通过 n/2k=1，我们可以求得 k=log2n，所以时间复杂度就是 O\(logn\)。

### 浅谈O\(logn\)

O\(logn\) 这种对数时间复杂度是一种极其高效的时间复杂度，有的时候甚至比时间复杂度是常量级 O\(1\) 的算法还要高效。为什么这么说呢？

因为 logn 是一个非常“恐怖”的数量级，即便 n 非常非常大，对应的 logn 也很小。比如 n 等于 2 的 32 次方，这个数很大了吧？大约是 42 亿。也就是说，如果我们在 42 亿个数据中用二分查找一个数据，最多需要比较 32 次。

我们前面讲过，用大 O 标记法表示时间复杂度的时候，会省略掉常数、系数和低阶。对于常量级时间复杂度的算法来说，O\(1\) 有可能表示的是一个非常大的常量值，比如 O\(1000\)、O\(10000\)。所以，常量级时间复杂度的算法有时候可能还没有 O\(logn\) 的算法执行效率高。

## 二分查找代码实现

### 递归方式

```go
func BsearchRecu(array []int, n int, value int) int {
	return bsearch(array, 0, n - 1, value)
}

func bsearch(array []int, low int, high int, value int) int {
	if low > high{
		return -1
	}

	mid := low + ((high - low) >> 1)
	if array[mid] == value{
		return mid
	}else if array[mid] < value{
		return bsearch(array, mid + 1, high, value)
	}else{
		return bsearch(array, low, mid - 1, value)
	}
}
```

### 非递归方式

```go
func BSearch(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] == value{
			return mid
		}else if array[mid] < value{
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return -1
}
```

## 二分查找应用场景的局限性

### 二分查找依赖的是顺序表结构

二分查找只能应用于数组，那二分查找能否依赖其他数据结构呢？比如链表。答案是不可以的，主要原因是二分查找算法需要按照下标随机访问元素。数组按照下标随机访问数据的时间复杂度是 O\(1\)，而链表随机访问的时间复杂度是 O\(n\)。所以，如果数据使用链表存储，二分查找的时间复杂就会变得很高。

### 二分查找针对的是有序数据

二分查找对这一点的要求比较苛刻，数据必须是有序的。如果数据没有序，我们需要先排序。排序的时间复杂度最低是 O\(nlogn\)。所以，如果我们针对的是一组静态的数据，没有频繁地插入、删除，我们可以进行一次排序，多次二分查找。这样排序的成本可被均摊，二分查找的边际成本就会比较低。

但是，如果我们的数据集合有频繁的插入和删除操作，要想用二分查找，要么每次插入、删除操作之后保证数据仍然有序，要么在每次二分查找之前都先进行排序。针对这种动态数据集合，无论哪种方法，维护有序的成本都是很高的。

所以，二分查找只能用在插入、删除操作不频繁，一次排序多次查找的场景中。针对动态变化的数据集合，二分查找将不再适用。

### 数据量太小不适合二分查找

如果要处理的数据量很小，完全没有必要用二分查找，顺序遍历就足够了。比如我们在一个大小为 10 的数组中查找一个元素，不管用二分查找还是顺序遍历，查找速度都差不多。只有数据量比较大的时候，二分查找的优势才会比较明显。

不过，这里有一个例外。如果数据之间的比较操作非常耗时，不管数据量大小，都推荐使用二分查找。比如，数组中存储的都是长度超过 300 的字符串，如此长的两个字符串之间比对大小，就会非常耗时。我们需要尽可能地减少比较次数，而比较次数的减少会大大提高性能，这个时候二分查找就比顺序遍历更有优势。

### 数据量太大也不适合二分查找

二分查找的底层需要依赖数组这种数据结构，而数组为了支持随机访问的特性，要求内存空间连续，对内存的要求比较苛刻。比如，我们有 1GB 大小的数据，如果希望用数组来存储，那就需要 1GB 的连续内存空间。

注意这里的“连续”二字，也就是说，即便有 2GB 的内存空间剩余，但是如果这剩余的 2GB 内存空间都是零散的，没有连续的 1GB 大小的内存空间，那照样无法申请一个 1GB 大小的数组。而我们的二分查找是作用在数组这种数据结构之上的，所以太大的数据用数组存储就比较吃力了，也就不能用二分查找了。

## 二分查找变形问题

### 查找第一个值等于给定值的元素

如何在有重复的数据中找到第一个值等于给定值的数据。

比如下面这样一个有序数组，其中，a\[5\]，a\[6\]，a\[7\]的值都等于 8，是重复的数据。我们希望查找第一个等于 8 的数据，也就是下标是 5 的元素。

![](../.gitbook/assets/image%20%28108%29.png)

如果我们用之前实现的二分查找那在a\[7\]时正好等于8就返回了，但是a\[7\]并不是8出现的第一次，所以并不符合之前的需求。

#### 代码实现

```go
func BSearchFirst(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] < value{
			low = mid + 1
		}else if array[mid] > value{
			high = mid - 1
		}else{
			if mid == 0 || array[mid - 1] != value{
				return mid
			}else {
				high = mid - 1
			}
		}
	}
	return -1
}
```

a\[mid\]跟要查找的 value 的大小关系有三种情况：大于、小于、等于。

* 对于 a\[mid\]&gt;value 的情况，我们需要更新 high= mid-1；
* 对于 a\[mid\]&lt;value 的情况，我们需要更新 low=mid+1；
* 对于 a\[mid\]==value 的情况
  * 如果 mid 等于 0，那这个元素已经是数组的第一个元素，那它肯定是我们要找的；
  * 如果 mid 不等于 0，但 a\[mid\]的前一个元素 a\[mid-1\]不等于 value，那也说明 a\[mid\]就是我们要找的第一个值等于给定值的元素。
  * 如果a\[mid\]前面的一个元素 a\[mid-1\]也等于 value，那说明此时的 a\[mid\]肯定不是我们要查找的第一个值等于给定值的元素。那我们就更新 high=mid-1，因为要找的元素肯定出现在\[low, mid-1\]之间。

### 查找最后一个值等于给定值的元素

如何在有重复的数据中找到最后一个值等于给定值的数据。

比如下面这样一个有序数组，其中，a\[5\]，a\[6\]，a\[7\]的值都等于 8，是重复的数据。我们希望查找最后一个等于 8 的数据，也就是下标是 7 的元素。

![](../.gitbook/assets/image%20%28108%29.png)

#### 代码实现

```go
func BSearchLast(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] < value{
			low = mid + 1
		}else if array[mid] > value{
			high = mid - 1
		}else{
			if mid == n - 1 || array[mid + 1] != value{
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}
```

a\[mid\]跟要查找的 value 的大小关系有三种情况：大于、小于、等于。

* 对于 a\[mid\]&gt;value 的情况，我们需要更新 high= mid-1；
* 对于 a\[mid\]&lt;value 的情况，我们需要更新 low=mid+1；
* 对于 a\[mid\]==value 的情况
  * 如果 mid 等于 n - 1，那这个元素已经是数组的最后一个元素，那它肯定是我们要找的；
  * 如果 mid 不等于 n - 1，但 a\[mid\]的后一个元素 a\[mid+1\]不等于 value，那也说明 a\[mid\]就是我们要找的第一个值等于给定值的元素。
  * 如果a\[mid\]后面的一个元素 a\[mid+1\]也等于 value，那说明此时的 a\[mid\]肯定不是我们要查找的最后一个值等于给定值的元素。那我们就更新 low=mid+1，因为要找的元素肯定出现在\[mid + 1, high\]之间。

### 查找第一个大于等于给定值的元素

在有序数组中，查找第一个大于等于给定值的元素。比如，数组中存储的这样一个序列：3，4，6，7，10。如果查找第一个大于等于 5 的元素，那就是 6。

#### 代码实现

```go
func BSearchFirstBigger(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] >= value{
			if mid == 0 || array[mid - 1] < value{
				return mid
			}else{
				high= mid - 1
			}
		}else {
			low = mid + 1
		}
	}
	return -1
}
```

如果 a\[mid\]小于要查找的值 value，那要查找的值肯定在\[mid+1, high\]之间，所以，我们更新 low=mid+1。

对于 a\[mid\]大于等于给定值 value 的情况，我们要先看下这个 a\[mid\]是不是我们要找的第一个值大于等于给定值的元素。如果 a\[mid\]前面已经没有元素，或者前面一个元素小于要查找的值 value，那 a\[mid\]就是我们要找的元素。

如果 a\[mid-1\]也大于等于要查找的值 value，那说明要查找的元素在\[low, mid-1\]之间，所以，我们将 high 更新为 mid-1。

### 查找最后一个小于等于给定值的元素

比如，数组中存储了这样一组数据：3，5，6，8，9，10。最后一个小于等于 7 的元素就是 6。

#### 代码实现

```go
func BSearchLastSmaller(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] > value{
			high = mid -1
		}else {
			if mid == n - 1 || array[mid + 1] > value{
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}
```

如果 a\[mid\]大于要查找的值 value，那要查找的值肯定在\[low, mid-1\]之间，所以，我们更新 high=mid-1。

对于 a\[mid\]大于等于给定值 value 的情况，我们要先看下这个 a\[mid\]是不是我们要找的最后一个值小于等于给定值的元素。如果 a\[mid\]后面已经没有元素，或者后面一个元素大于要查找的值 value，那 a\[mid\]就是我们要找的元素。

如果 a\[mid+1\]也大于等于要查找的值 value，那说明要查找的元素在\[mid + 1, high\]之间，所以，我们将 low 更新为 mid+1。

