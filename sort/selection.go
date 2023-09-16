/**
 * @Author: Hardews
 * @Date: 2023/9/16 15:55
 * @Description: 选择排序
**/

package sort

func SelectionSort(target []int) {
	for i := range target {
		// 记录下标
		var min = i
		for j := i; j < len(target); j++ {
			// 当前元素更小，交换下标
			if target[j] < target[min] {
				min = j
			}
		}
		// 最小的下标与当前下标交换
		temp := target[min]
		target[min] = target[i]
		target[i] = temp
	}
}
