/**
 * @Author: Hardews
 * @Date: 2023/9/16 16:09
 * @Description: 插入排序
**/

package sort

func InsertionSort(target []int) {
	for i := range target {
		// 前一个元素
		preIndex := i - 1
		current := target[i]
		// 从小到大排序
		for preIndex >= 0 && current < target[preIndex] {
			// 前后俩元素交换，直到找到要插入的 index
			target[preIndex+1] = target[preIndex]
			preIndex--
		}
		target[preIndex+1] = current
	}
}
