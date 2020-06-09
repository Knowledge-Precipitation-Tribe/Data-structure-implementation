# 基数排序

基数排序是按照低位先排序，然后收集；再按照高位排序，然后再收集；依次类推，直到最高位。有时候有些属性是有优先级顺序的，先按低优先级排序，再按高优先级排序。最后的次序就是高优先级高的在前，高优先级相同的低优先级高的在前。

![](../../.gitbook/assets/image%20%2873%29.png)

基数排序对要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系，如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。除此之外，每一位的数据范围不能太大，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O\(n\) 了。

## **算法描述**

* 取得数组中的最大数，并取得位数；
* arr为原始数组，从最低位开始取每个位组成radix数组；
* 对radix进行计数排序（利用计数排序适用于小范围数的特点）；

## **动图演示**

![](../../.gitbook/assets/radix-sort.gif)

## 代码实现

```go
func RadixSort(arr []int, maxDigit int) []int{
	bucket := make([][]int, maxDigit + 1)
	length := len(arr)
	for i := 0; i<maxDigit + 1; i++{
		bucket[i] = make([]int, 0)
	}

	for i :=0; i <length; i++{
		index := arr[i] % 10
		bucket[index] = append(bucket[index], arr[i])
	}

	for i := 0; i<maxDigit + 1; i++{
		if len(bucket[i]) > 1{
			bucket[i] = BubbleSort(bucket[i])
		}
	}

	result := make([]int, 0)
	for i := 0; i<maxDigit + 1; i++{
		bucket[i] = BubbleSort(bucket[i])
		for j := 0; j<len(bucket[i]); j++{
			result = append(result, bucket[i][j])
		}
	}
	return result
}
```

## 时间复杂度分析

基数排序基于分别排序，分别收集，所以是稳定的。但基数排序的性能比桶排序要略差，每一次关键字的桶分配都需要O\(n\)的时间复杂度，而且分配之后得到新的关键字序列又需要O\(n\)的时间复杂度。假如待排数据可以分为d个关键字，则基数排序的时间复杂度将是O\(d\*2n\) ，当然d要远远小于n，因此基本上还是线性级别的。

基数排序的空间复杂度为O\(n+k\)，其中k为桶的数量。一般来说n&gt;&gt;k，因此额外空间需要大概n个左右。

