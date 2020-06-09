package sort_algo

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
