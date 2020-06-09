package sort_algo

func CountingSort(arr []int, maxValue int)[]int{
	bucket := make([]int, maxValue+1)
	length := len(arr)

	for i := 0; i<length; i++{
		bucket[arr[i]]++
	}

	i := 0
	for j := 0; j<maxValue + 1; j++{
		for bucket[j] > 0{
			arr[i] = j
			i++
			bucket[j]--
		}
	}
	return arr
}