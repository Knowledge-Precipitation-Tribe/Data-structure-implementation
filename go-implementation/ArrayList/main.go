package main

import (
	"fmt"
	"go-implementation/ArrayList/Array"
)

func main() {
	arrayList := &Array.ArrayList{}
	arrayList.Init()
	arrayList.AppendElem(1)
	arrayList.AppendElem(2)
	arrayList.AppendElem(3)
	arrayList.InsertElemFast(0, -1)

	fmt.Println(arrayList.FindElem(3))
	fmt.Println(arrayList.GetElem(2))

	arrayList.ShowList()

	arrayList.DeleteElem(3)
	arrayList.DeleteElemByIndex(0)
	arrayList.ShowList()
}