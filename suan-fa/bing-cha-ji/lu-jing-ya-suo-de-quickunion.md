# 路径压缩的quick-union

```text
type UF struct {
	id []int
	sz []int
	count int
}

func (uf *UF) Init(n int){
	uf.count = n
	uf.id = make([]int, n)
	uf.sz = make([]int, n)
	for i := 0; i < n; i++{
		uf.id[i] = i
	}
	for i := 0; i < n; i++{
		uf.sz[i] = i
	}
}

//返回元素中连通分量的个数
func (uf *UF) Count() int {
	return uf.count
}

//实现方式4：采用路径压缩的方法，让每个结点的直接上级就是根结点，构造一颗扁平的树
//查询一个变量所属的连通分量
func (uf *UF) Find(p int) int {
	if p != uf.id[p]{
		uf.id[p] = uf.Find(uf.id[p])
	}
	return uf.id[p]
}

//在p和q之间添加连接
func (uf *UF) Union(p int, q int){
	pRoot := uf.Find(p)
	qRoot := uf.Find(q)

	if pRoot == qRoot {
		return
	}
	if uf.sz[pRoot] < uf.sz[qRoot]{
		uf.id[pRoot] = qRoot
		uf.sz[qRoot] += uf.sz[pRoot]
	}else{
		uf.id[qRoot] = pRoot
		uf.sz[pRoot] += uf.sz[qRoot]
	}

	uf.count--
}


//返回两个变量是否连通
func (uf *UF) Connected(p int, q int)bool{
	return uf.Find(p) == uf.Find(q)
}
```

