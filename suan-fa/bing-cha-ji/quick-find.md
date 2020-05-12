# quick-find

```text
type UF struct {
	id []int
	count int
}

func (uf *UF) Init(n int){
	uf.count = n
	uf.id = make([]int, n)
	for i := 0; i < n; i++{
		uf.id[i] = i
	}
}

//返回元素中连通分量的个数
func (uf *UF) Count() int {
	return uf.count
}

//实现方式1：判断id[p] == id[q],如果两个所在的是不同连通分量，
//要将p的所有连通分量值进行修改
//但是这种方式显然不行，每次操作都要扫描整个id
//查询一个变量所属的连通分量
func (uf *UF) Find(p int) int {
	return uf.id[p]
}

//在p和q之间添加连接
func (uf *UF) Union(p int, q int){
	pID := uf.Find(p)
	qID := uf.Find(q)

	if pID == qID {
		return
	}
	for i := 0; i < len(uf.id); i++{
		if uf.id[i] == pID{
			uf.id[i] = qID
		}
	}
	uf.count--
}


//返回两个变量是否连通
func (uf *UF) Connected(p int, q int)bool{
	return uf.Find(p) == uf.Find(q)
}
```

