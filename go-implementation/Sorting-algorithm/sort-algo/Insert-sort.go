package sort_algo

// 插入排序（Insertion-Sort）的算法描述是一种简单直观的排序算法。
// 它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
//
//算法描述
// 一般来说，插入排序都采用in-place在数组上实现。
// 具体算法描述如下：
// 从第一个元素开始，该元素可以认为已经被排序；
// 取出下一个元素，在已经排序的元素序列中从后向前扫描；
// 如果该元素（已排序）大于新元素，将该元素移到下一位置；
// 重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
// 将新元素插入到该位置后；
// 重复步骤2~5。
func InsertionSort(arr []int) []int {
	length := len(arr)
	var preIndex int
	var current int
	for i := 1; i < length; i++ {
		preIndex = i - 1
		current = arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex = preIndex - 1
		}
		arr[preIndex+1] = current
	}
	return arr
}
