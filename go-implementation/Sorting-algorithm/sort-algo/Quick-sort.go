package sort_algo

func QuickSort(arr []int, low, high int) {
	var temp int
	i := low
	j := high
	if low < high {
		temp = arr[low]
		//下面循环完成一次排序，即将数组中小于temp的关键字放在左边，大于temp的关键字放在右边
		for i < j {
			for j > i && arr[j] >= temp {
				j = j - 1
			}
			if i < j {
				arr[i] = arr[j]
				i = i + 1
			}
			for i < j && arr[i] < temp {
				i = i + 1
			}
			if i < j {
				arr[j] = arr[i]
				j = j - 1
			}
		}
		arr[i] = temp
		QuickSort(arr, low, i-1)
		QuickSort(arr, i+1, high)
	}
}
