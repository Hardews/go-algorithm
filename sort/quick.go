/**
 * @Author: Hardews
 * @Date: 2023/9/16 18:11
 * @Description:
**/

package sort

func QuickSort(target []int) []int {
	return quickSort(target, 0, len(target)-1)
}

func quickSort(target []int, left, right int) []int {
	if left < right {
		// 跟树的遍历差不多
		mid := partition(target, left, right)
		// 这里是 左边 到 中间值 -1（因为左右两边是围绕中间值排序的，中间值不用再排了
		quickSort(target, left, mid-1)
		// 中间值 + 1 到 右边
		quickSort(target, mid+1, right)
	}
	return target
}

func partition(target []int, left, right int) int {
	pivot := left
	pivotNum := target[pivot]
	for left < right {
		if target[right] >= pivotNum {
			right--
			continue
		}
		if target[left] <= pivotNum {
			left++
			continue
		}

		// 到这里证明它们要交换了
		swap(target, left, right)
	}
	// 循环结束， left == right
	swap(target, left, pivot)

	// 返回中间的值
	return left
}
