package set

import (
	"fmt"
	"testing"
)

/**
* @Author: super
* @Date: 2020-08-24 14:50
* @Description:
**/

var set = InitSet()

func TestMySet_AddElementSet(t *testing.T) {
	set.AddElementSet(1)
	set.AddElementSet(2)
	set.ShowSet()
}

func TestMySet_IsExistSet(t *testing.T) {
	fmt.Println(set.IsExistSet(1))
}

func TestMySet_DeleteElementSet(t *testing.T) {
	fmt.Println(set.DeleteElementSet(1))
	fmt.Println(set.DeleteElementSet(5))
	set.ShowSet()
}