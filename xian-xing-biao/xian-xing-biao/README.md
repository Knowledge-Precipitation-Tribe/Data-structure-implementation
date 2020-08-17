# 线性表

## 线性表的定义

```go
type ArrayList struct {
	data   []int //存储数据
	length int //线性表长度
}
```

## 常用操作

### 查找指定元素

```go
//找到元素返回下标, true，未找到返回-1, false
func (list *ArrayList) FindElem(x int) (int, bool) {
	for i := 0; i < list.length; i++ {
		if list.data[i] == x {
			return i, true
		}
	}
	return -1, false
}
```

### 查找指定位置的元素

```go
//获取指定位置的元素
func (list *ArrayList) GetElem(p int) (int, bool) {
	if p < 0 || p >= list.length{
		return -1, false
	}
	return list.data[p], true
}
```

### 快速插入元素

```go
//在线性表p位置插入元素e
//将原位置元素先添加到线性表末尾
//将新的元素插入的p的位置
func (list *ArrayList) InsertElemFast(p int, e int) bool {
	if p < 0 || p >= list.length {
		return false
	}
	if p == list.length {
		list.data = append(list.data, e)
	} else {
		temp := list.data[p]
		list.data = append(list.data, temp)
		list.data[p] = e
	}
	list.length = list.length + 1
	return true
}
```

### 插入元素

```go
//正常插入元素，将p位置元素依次向后移
func (list *ArrayList) InsertElem(p int, e int) bool {
	if p < 0 || p >= list.length {
		return false
	}
	for i := list.length - 1; i >= p; i-- {
		list.data[i+1] = list.data[i]
	}
	list.data[p] = e
	list.length = list.length + 1
	return true
}
```

### 添加元素

```go
func (list *ArrayList) AppendElem(e int){
	list.data = append(list.data, e)
	list.length = list.length + 1
}
```

### 删除指定元素

```go
//删除指定元素
func (list *ArrayList) DeleteElem(e int) (int, bool) {
	i, v := list.FindElem(e)
	if v != false{
		return list.DeleteElemByIndex(i)
	}else{
		return -1, false
	}
}
```

### 删除指定位置的元素

```go
//删除指定位置的元素
func (list *ArrayList) DeleteElemByIndex(p int) (int, bool) {
	if p < 0 || p >= list.length {
		return -1, false
	}
	e := list.data[p]
	for i := p; i < list.length-1; i++ {
		list.data[i] = list.data[i+1]
	}
	list.length = list.length - 1
	return e, true
}
```

### 显示顺序表内容

```go
func (list *ArrayList) ShowList() {
	for i := 0; i< list.length; i++{
		fmt.Printf("%d ", list.data[i])
	}
	fmt.Println()
}
```

