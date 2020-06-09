package main

import (
	"fmt"
	"go-implementation/Sorting-algorithm/sort-algo"
)

func main() {
	arr := []int{5,8,7,9,10,3,2,1}
	//arr = sort_algo.BubbleSort(arr)
	//arr = sort_algo.SelectSort(arr)
	arr = sort_algo.InsertionSort(arr)
	fmt.Println(arr)
}