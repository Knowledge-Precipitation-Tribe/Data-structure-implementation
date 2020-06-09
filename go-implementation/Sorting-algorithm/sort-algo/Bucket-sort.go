package sort_algo



func BucketSort(arr []int, bucketSize int) []int{
	length := len(arr)
	if length == 0{
		return arr
	}
	minValue := arr[0]
	maxValue := arr[0]

	for i := 1; i<length; i++{
		if arr[i] < minValue{
			minValue = arr[i]
		}else if arr[i] > maxValue{
			maxValue = arr[i]
		}
	}

	bucketCount := (maxValue - minValue) / bucketSize + 1
	buckets := make([][]int, bucketCount)
	for i := 0; i<bucketCount; i++{
		buckets[i] = make([]int, 0)
	}


	for i := 0; i<length; i++{
		index := (arr[i]-minValue)/bucketSize
		buckets[index] = append(buckets[index], arr[i])
	}

	result := make([]int, 0)
	for i := 0; i<bucketCount; i++{
		buckets[i] = BubbleSort(buckets[i])
		for j := 0; j<len(buckets[i]); j++{
			result = append(result, buckets[i][j])
		}
	}
	return result
}