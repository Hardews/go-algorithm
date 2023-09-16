/**
 * @Author: Hardews
 * @Date: 2023/9/16 16:45
 * @Description:
**/

package sort

func ShellsSort(target []int) {
	var gap int = len(target) / 2
	for gap > 0 {
		for i := gap; i < len(target); i++ {
			// 从第一个 gap 开始遍历
			// 上一个元素为 i - gap
			preIndex := i - gap
			current := target[i]
			for preIndex >= 0 && current < target[preIndex] {
				// 插入排序
				target[preIndex+gap] = target[preIndex]
				preIndex -= gap
			}
			target[preIndex+gap] = current
		}
		gap /= 2
	}
}
