/**
 * @Author: Hardews
 * @Date: 2023/9/16 15:38
 * @Description: 冒泡排序
**/

package sort

func BubbleSort(target []int) {
	for i := range target {
		// 遍历到 len(target) - i - 1 即可
		for j := 0; j < len(target)-i-1; j++ {
			if target[j] > target[j+1] {
				// 满足条件，交换
				temp := target[j]
				target[j] = target[j+1]
				target[j+1] = temp
			}
		}
	}
}
