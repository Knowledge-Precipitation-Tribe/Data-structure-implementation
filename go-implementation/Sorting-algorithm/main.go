package main

import (
	"fmt"
	"go-implementation/Sorting-algorithm/sort-algo"
)

func main() {
	arr := []int{5,8,7,9,10,3,2,1}
	arr = sort_algo.BubbleSort(arr)
	fmt.Println(arr)
}