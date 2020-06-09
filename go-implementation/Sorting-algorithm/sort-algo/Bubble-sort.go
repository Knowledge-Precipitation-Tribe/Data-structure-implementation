package sort_algo

func BubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i< length - 1; i++{
		for j := 0; j < length - i - 1; j++{
			if arr[j] > arr[j + 1]{
				temp := arr[j + 1]
				arr[j + 1] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}