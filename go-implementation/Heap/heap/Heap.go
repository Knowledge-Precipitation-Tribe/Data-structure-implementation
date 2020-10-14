package heap

import "errors"

type Heap struct {
	data  []int
	n     int
	count int
}

func NewHeap(capacity int) *Heap {
	heap := &Heap{
		data:  make([]int, capacity),
		n:     capacity,
		count: 0,
	}
	return heap
}

func (h Heap) Insert(data int) error{
	if h.count >= h.n{
		return errors.New("堆已满")
	}
	h.data[h.count] = data
	i := h.count
	for i/2 > 0 && h.data[i] > h.data[i/2]{
		swap(h.data,i,i/2)
		i = i /2
	}
	return nil
}

func swap(data []int, i int, j int) {
	temp := data[i]
	data[i] = data[j]
	data[j] = temp
}