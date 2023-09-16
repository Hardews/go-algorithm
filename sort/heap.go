/**
 * @Author: Hardews
 * @Date: 2023/9/16 19:42
 * @Description:
**/

package sort

func HeapSort(target []int) {
	// 创建最大堆
	shiftDown(target)
	arrLen := len(target)
	for i := len(target) - 1; i >= 0; i-- {
		swap(target, 0, i)
		arrLen -= 1
		heapify(target, 0, arrLen)
	}
}

// shiftDown 创建最大堆
func shiftDown(target []int) {
	for i := len(target) / 2; i >= 0; i-- {
		heapify(target, i, len(target))
	}
}

// heapify 以传入参数为堆顶，调整最大堆
func heapify(target []int, start, end int) {
	var largest = start
	var left = 2*start + 1  // 左子节点
	var right = 2*start + 2 // 右子节点
	// 比较左节点
	if left < end && target[left] > target[largest] {
		largest = left
	}

	// 比较右节点
	if right < end && target[right] > target[largest] {
		largest = right
	}

	if largest != start {
		// 堆顶不是最大值了，下沉
		swap(target, start, largest)
		heapify(target, largest, end)
	}
}
