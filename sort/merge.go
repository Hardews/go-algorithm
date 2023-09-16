/**
 * @Author: Hardews
 * @Date: 2023/9/16 17:15
 * @Description:
**/

package sort

func MergeSort(target []int) []int {
	if len(target) < 2 {
		// 只剩一个元素或者没元素了就可以返回
		return target
	}
	mid := len(target) / 2
	left := target[:mid]
	right := target[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left, right []int) []int {
	var res []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	if len(left) == 0 {
		res = append(res, right...)
	}

	if len(right) == 0 {
		res = append(res, left...)
	}

	return res
}
