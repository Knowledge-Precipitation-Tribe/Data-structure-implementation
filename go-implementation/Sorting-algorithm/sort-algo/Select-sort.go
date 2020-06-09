package sort_algo

func SelectSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j<length; j++{
			if arr[j] < arr[minIndex]{
				minIndex = j
			}
		}
		if minIndex != i{
			temp := arr[i]
			arr[i] = arr[minIndex]
			arr[minIndex] = temp
		}
	}
	return arr
}
