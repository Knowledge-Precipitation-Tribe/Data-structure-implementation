package algo

//非递归实现二分查找
func BSearch(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] == value{
			return mid
		}else if array[mid] < value{
			low = mid + 1
		}else{
			high = mid - 1
		}
	}
	return -1
}


//递归实现二分查找
func BsearchRecu(array []int, n int, value int) int {
	return bsearch(array, 0, n - 1, value)
}

func bsearch(array []int, low int, high int, value int) int {
	if low > high{
		return -1
	}

	mid := low + ((high - low) >> 1)
	if array[mid] == value{
		return mid
	}else if array[mid] < value{
		return bsearch(array, mid + 1, high, value)
	}else{
		return bsearch(array, low, mid - 1, value)
	}
}

//查找第一个值等于给定值的元素
func BSearchFirst(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] < value{
			low = mid + 1
		}else if array[mid] > value{
			high = mid - 1
		}else{
			if mid == 0 || array[mid - 1] != value{
				return mid
			}else {
				high = mid - 1
			}
		}
	}
	return -1
}

//查找第一个值等于给定值的元素
func BSearchLast(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] < value{
			low = mid + 1
		}else if array[mid] > value{
			high = mid - 1
		}else{
			if mid == n - 1 || array[mid + 1] != value{
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}

//查找第一个大于等于给定值的元素
func BSearchFirstBigger(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] >= value{
			if mid == 0 || array[mid - 1] < value{
				return mid
			}else{
				high= mid - 1
			}
		}else {
			low = mid + 1
		}
	}
	return -1
}

//查找最后一个小于等于给定值的元素
func BSearchLastSmaller(array []int, n int, value int) int {
	low := 0
	high := n - 1

	for low <= high{
		mid := (low + high) / 2
		if array[mid] > value{
			high = mid -1
		}else {
			if mid == n - 1 || array[mid + 1] > value{
				return mid
			}else {
				low = mid + 1
			}
		}
	}
	return -1
}