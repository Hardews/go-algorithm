/**
 * @Author: Hardews
 * @Date: 2023/9/17 1:09
 * @Description:
**/

package sort

func RadixSort(target []int) {
	// 找到最大位
	var max int
	for _, num := range target {
		if num > max {
			max = num
		}
	}

	var maxBit int
	for max > 0 {
		max /= 10
		maxBit++
	}

	for i := 1; i <= maxBit; i++ {
		bitSort(target, 10, i)
	}
}

// bitSort bit 当前位，个位是 1，scale 进制，10 进制为 10
func bitSort(target []int, scale, bit int) {
	// 初始化一个桶
	var buket = make([][]int, scale)
	for _, num := range target {
		// 获取当前位的数字
		bitDiv := (bit - 1) * 10
		if bitDiv == 0 {
			bitDiv = 1
		}
		bitNum := (num / bitDiv) % 10
		// 放入桶
		buket[bitNum] = append(buket[bitNum], num)
	}

	// 将桶中的元素依次放进 target 中
	var i int
	for j := range buket {
		for _, num := range buket[j] {
			target[i] = num
			i++
		}
	}
}
