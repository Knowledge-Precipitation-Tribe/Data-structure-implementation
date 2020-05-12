# 加权quick-union

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

//实现方式3：采用树的形式，找到当前结点所属的根结点，判断两个变量是否属于统一连通分量,
// 但是两树合并时采用将小树往大树的身上合并
//查询一个变量所属的连通分量
func (uf *UF) Find(p int) int {
	for p != uf.id[p]{
		p = uf.id[p]
	}
	return p
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

