package main

import "fmt"

/**
* @Author: super
* @Date: 2021-03-06 15:57
* @Description:
**/
func twoSum(nums []int, target int) []int {
	hashMap := map[int]int{}
	for i, x := range nums {
		if p, ok := hashMap[target-x]; ok {
			return []int{p, i}
		}
		hashMap[x] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	//nums := []int{}
	fmt.Println(twoSum(nums, 9))
}