package main

import (
	"fmt"
	"go-implementation/Search/algo"
)

func main() {
	array := []int{3,4,6,7,10}
	//fmt.Println(algo.BsearchRecu(array, 9, 4))
	//fmt.Println(algo.BSearch(array, 9, 4))
	//fmt.Println(algo.BSearchFirst(array, 10, 8))
	//fmt.Println(algo.BSearchLast(array, 10, 8))
	fmt.Println(algo.BSearchFirstBigger(array, 5, 5))
}