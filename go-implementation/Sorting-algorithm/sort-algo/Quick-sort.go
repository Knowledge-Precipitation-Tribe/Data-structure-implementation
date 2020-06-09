package sort_algo

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}


func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partition(a, lo, hi)
	QuickSort(a, lo, p-1)
	QuickSort(a, p+1, hi)
}


// 并发版
func QuickSortConcurrent(a []int, lo, hi int, done chan struct{}, depth int) {
	if lo >= hi {
		done <- struct{}{}
		return
	}
	depth--
	p := partition(a, lo, hi)
	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go QuickSortConcurrent(a, lo, p-1, childDone, depth)
		go QuickSortConcurrent(a, p+1, hi, childDone, depth)
		<-childDone
		<-childDone
	} else {
		QuickSort(a, lo, p-1)
		QuickSort(a, p+1, hi)
	}
	done <- struct{}{}
}