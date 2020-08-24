package set

import "fmt"

/**
* @Author: super
* @Date: 2020-08-24 14:44
* @Description:
**/

type MySet struct {
	data   map[interface{}]bool
	length int
}

func InitSet() *MySet {
	return &MySet{
		data:   make(map[interface{}]bool),
		length: 0,
	}
}

func (set *MySet) AddElementSet(v interface{}) {
	if !set.data[v] {
		set.data[v] = true
		set.length++
	}
}

func (set *MySet) IsExistSet(v interface{}) bool {
	return set.data[v]
}

//return true 删除成功
//return false 删除失败
func (set *MySet) DeleteElementSet(v interface{}) bool {
	if set.data[v] {
		delete(set.data, v)
		return true
	}
	return false
}

func (set *MySet) ShowSet() {
	for k := range set.data{
		fmt.Println(k)
	}
}
