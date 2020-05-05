package main

import (
	"fmt"
	unionFind "go-implementation/union-find/weighted-quick-union"
)

func main() {
	n := 10
	uf := unionFind.UF{}
	uf.Init(n)
	for i := 0; i < 11; i++{
		var p int
		var q int
		fmt.Scanf("%d", &p)
		fmt.Scanf("%d", &q)
		if uf.Connected(p, q){
			fmt.Printf("%d and %d is connected\n", p, q)
			continue
		}
		uf.Union(p, q)
		fmt.Printf("%d and %d union\n", p ,q)
	}
	fmt.Println("count: ", uf.Count())
	fmt.Println(uf.GetID())
}