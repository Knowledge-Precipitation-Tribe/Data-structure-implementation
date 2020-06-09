package sort_algo

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