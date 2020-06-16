package main

import (
	"fmt"
	"go-implementation/Search/algo"
)

func main() {
	array := []int{1,3,4,5,6,8,8,8,11,18}
	//fmt.Println(algo.BsearchRecu(array, 9, 4))
	//fmt.Println(algo.BSearch(array, 9, 4))
	fmt.Println(algo.BSearchFirst(array, 10, 8))
}