package sort_algo

import "fmt"

/**
* @Author: super
* @Date: 2021-03-06 15:25
* @Description:
**/

// 堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。、
// 堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：
// 即子结点的键值或索引总是小于（或者大于）它的父节点。
//
//算法描述：
// 将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
// 将堆顶元素R[1]与最后一个元素R[n]交换，
// 此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn)，且满足R[1,2…n-1]<=R[n]；
// 由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，
// 然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。
// 不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
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
