package main

import (
	"fmt"
	sort_algo "go-implementation/Sorting-algorithm/sort-algo"
)

func main() {
	arr := []int{5, 8, 7, 9, 10, 3, 2, 1, 4}
	//arr = sort_algo.BubbleSort(arr)
	//arr = sort_algo.SelectSort(arr)
	//arr = sort_algo.InsertionSort(arr)
	//sort_algo.MergeSort(arr, 0, len(arr) - 1)
	//sort_algo.QuickSort(arr, 0, len(arr) - 1)
	//arr = sort_algo.BucketSort(arr, 5)
	//arr = sort_algo.CountingSort(arr, 10)
	//arr = sort_algo.RadixSort(arr, 9)
	sort_algo.ShellSort(arr)
	//sort_algo.HeapSort(arr)
	fmt.Println(arr)
}
