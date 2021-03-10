package sort_algo

import "fmt"

/**
* @Author: super
* @Date: 2021-03-06 15:25
* @Description:
**/

func adjustHeap(a []int, pos int, lenght int) {
	for {
		// 计算孩子位置
		child := pos*2 + 1
		// 检查孩子是否越界
		if child >= (lenght - 1) {
			break
		}

		// 找出孩子中较大的那个
		if a[child+1] > a[child] {
			child++
		}
		// 检查孩子是否大于父亲，如果大于则交换
		if a[pos] < a[child] {
			// 父子交换
			a[pos], a[child] = a[child], a[pos]
			// 更新位置，指向子节点
			pos = child
		} else {
			// 如果子节点都小于父节点了，那就可以结束调整了
			break
		}
	}
}

func buildHeap(a []int) {
	// 从底层向上层构建，len(a)/2-1是第一个非叶子
	for i := len(a)/2 - 1; i >= 0; i-- {
		adjustHeap(a, i, len(a))
	}
}

func HeapSort(a []int) {
	// 构建大顶堆
	buildHeap(a)

	fmt.Println("buildHeap over:", a)
	fmt.Println("===============================")

	// 首尾交换，重新构建大顶堆
	for i := len(a) - 1; i >= 0; i-- {
		a[0], a[i] = a[i], a[0]
		adjustHeap(a, 0, i)
	}

}

func main() {
	num := []int{98, 48, 777, 63, 57, 433, 23, 1112, 1}
	HeapSort(num)
	fmt.Println("heap sort over:", num)
}