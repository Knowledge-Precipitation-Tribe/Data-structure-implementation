package main

import (
	"fmt"
	"go-implementation/union-find/quick-find"
)

func main() {
	n := 10
	uf := quick_find.UF{}
	uf.Init(n)
	for i := 0; i < 8; i++{
		var p int
		var q int
		fmt.Scanf("%d", &p)
		fmt.Scanf("%d", &q)
		if uf.Connected(p, q){
			continue
		}
		uf.Union(p, q)
		fmt.Printf("%d and %d union\n", p ,q)
	}
	fmt.Println("count: ", uf.Count())
}