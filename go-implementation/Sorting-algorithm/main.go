package main

import (
	"fmt"
	"go-implementation/Sorting-algorithm/sort-algo"
)

func main() {
	arr := []int{5, 8, 7, 9, 10, 3, 2, 1, 4}
	//arr = sort_algo.BubbleSort(arr)
	//arr = sort_algo.SelectSort(arr)
	//arr = sort_algo.InsertionSort(arr)
	//sort_algo.MergeSort(arr, 0, len(arr) - 1)
	//sort_algo.QuickSort(arr, 0, len(arr) - 1)
	done := make(chan struct{})
	sort_algo.QuickSortConcurrent(arr, 0, len(arr) - 1, done, 0)
	fmt.Println(arr)
}
